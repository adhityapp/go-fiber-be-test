package middleware

import (
	"fmt"
	"time"

	"github.com/bangadam/go-fiber-starter/utils/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected() fiber.Handler {
	conf := config.NewConfig()

	if conf.Middleware.Jwt.Secret == "" {
		panic("JWT secret is not set")
	}

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(conf.Middleware.Jwt.Secret),
		ContextKey:   "jwt",
		TokenLookup:  "header:Authorization",
		AuthScheme:   "Bearer",
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {

	fmt.Println(err)

	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

type JWTClaims struct {
	Token string `json:"token"`
	Type  string `json:"type"`
	jwt.RegisteredClaims
}

func JWTMiddleware(c *fiber.Ctx) error {
	conf := config.NewConfig()

	mySigningKey := []byte(conf.Middleware.Jwt.Secret)
	tokenString := c.Get("Authorization")[7:] // Strip "Bearer " prefix

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user", claims["user"])

	return c.Next()
}

func GenerateTokenAccess(employee_code string) (*string, error) {
	conf := config.NewConfig()

	mySigningKey := []byte(conf.Middleware.Jwt.Secret)
	// Create the Claims
	exp := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"user": employee_code,
		"exp":  exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		return nil, err
	}

	return &ss, nil
}
