package main

import (
	"github.com/bangadam/go-fiber-starter/app/middleware"
	"github.com/bangadam/go-fiber-starter/app/module/attendance"
	"github.com/bangadam/go-fiber-starter/app/module/auth"
	"github.com/bangadam/go-fiber-starter/app/module/department"
	"github.com/bangadam/go-fiber-starter/app/module/employee"
	"github.com/bangadam/go-fiber-starter/app/module/location"
	"github.com/bangadam/go-fiber-starter/app/module/position"
	"github.com/bangadam/go-fiber-starter/app/module/report"
	"github.com/bangadam/go-fiber-starter/app/router"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/utils/config"
	fxzerolog "github.com/efectn/fx-zerolog"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

// @title                       Go Fiber Starter API Documentation
// @version                     1.0
// @description                 This is a sample API documentation.
// @termsOfService              http://swagger.io/terms/
// @contact.name                Developer
// @contact.email               bangadam.dev@gmail.com
// @license.name                Apache 2.0
// @license.url                 http://www.apache.org/licenses/LICENSE-2.0.html
// @host                        localhost:8080
// @schemes                     http https
// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description                 "Type 'Bearer {TOKEN}' to correctly set the API Key"
// @BasePath                    /
func main() {
	fx.New(
		/* provide patterns */
		// config
		fx.Provide(config.NewConfig),
		// logging
		fx.Provide(bootstrap.NewLogger),
		// fiber
		fx.Provide(bootstrap.NewFiber),
		// database
		fx.Provide(database.NewDatabase),
		// middleware
		fx.Provide(middleware.NewMiddleware),
		// router
		fx.Provide(router.NewRouter),

		// provide modules
		auth.NewAuthModule,
		attendance.NewAttendanceModule,
		employee.NewEmployeeModule,
		department.NewDepartmentModule,
		location.NewLocationModule,
		position.NewPositionModule,
		report.NewReportModule,

		// start aplication
		fx.Invoke(bootstrap.Start),

		// define logger
		fx.WithLogger(fxzerolog.Init()),
	).Run()

	// app := fiber.New()
	// app.Get("/ping", func(c *fiber.Ctx) error {
	// 	return c.SendString("pong")
	// })

	// err := app.Listen(":8081")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("server started")
}
