package service

import (
	"errors"

	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/middleware"
	"github.com/bangadam/go-fiber-starter/app/module/auth/request"
	"github.com/bangadam/go-fiber-starter/app/module/auth/response"
	user_repo "github.com/bangadam/go-fiber-starter/app/module/user/repository"
)

// AuthService
type articleService struct {
	userRepo user_repo.UserRepository
}

// define interface of IAuthService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . AuthService
type AuthService interface {
	Login(req request.LoginRequest) (res response.LoginResponse, err error)
	Register(req request.RegisterRequest) (res response.RegisterResponse, err error)
}

// init AuthService
func NewAuthService(userRepo user_repo.UserRepository) AuthService {
	return &articleService{
		userRepo: userRepo,
	}
}

func (_i *articleService) Login(req request.LoginRequest) (res response.LoginResponse, err error) {
	// check user by email
	user, err := _i.userRepo.FindUserByEmployeeCode(req.EmployeeCode)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("user not found")
		return
	}

	// check password
	if !user.ComparePassword(req.Password) {
		err = errors.New("password not match")
		return
	}

	// do create token
	claims, err := middleware.GenerateTokenAccess(req.EmployeeCode)
	if err != nil {
		return
	}

	res.Token = claims

	return
}

func (_i *articleService) Register(req request.RegisterRequest) (res response.RegisterResponse, err error) {
	// check user by email
	user, err := _i.userRepo.FindUserByEmployeeCode(req.EmployeeCode)
	if err != nil {
		return
	}

	if user != nil {
		err = errors.New("email already exists")
		return
	}

	// do create user
	newUser := &schema.Employee{
		EmployeeCode: req.EmployeeCode,
		Password:     req.Password,
	}

	user, err = _i.userRepo.CreateUser(newUser)
	if err != nil {
		return
	}

	res.EmployeeID = user.EmployeeID
	res.EmployeeCode = user.EmployeeCode

	return
}
