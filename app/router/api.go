package router

import (
	"github.com/bangadam/go-fiber-starter/app/middleware"
	"github.com/bangadam/go-fiber-starter/app/module/attendance"
	"github.com/bangadam/go-fiber-starter/app/module/auth"
	"github.com/bangadam/go-fiber-starter/app/module/department"
	"github.com/bangadam/go-fiber-starter/app/module/employee"
	"github.com/bangadam/go-fiber-starter/app/module/location"
	"github.com/bangadam/go-fiber-starter/app/module/position"
	"github.com/bangadam/go-fiber-starter/app/module/report"
	"github.com/bangadam/go-fiber-starter/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct {
	App fiber.Router
	Cfg *config.Config

	AuthRouter       *auth.AuthRouter
	AttendanceRouter *attendance.AttendanceRouter
	EmployeeRouter   *employee.EmployeeRouter
	DepartmentRouter *department.DepartmentRouter
	LocationRouter   *location.LocationRouter
	PositionRouter   *position.PositionRouter
	ReportRouter     *report.ReportRouter
}

func NewRouter(
	fiber *fiber.App,
	cfg *config.Config,

	authRouter *auth.AuthRouter,
	attendanceRouter *attendance.AttendanceRouter,
	employeeRouter *employee.EmployeeRouter,
	departmentRouter *department.DepartmentRouter,
	locationRouter *location.LocationRouter,
	positionRouter *position.PositionRouter,
	reportRouter *report.ReportRouter,
) *Router {
	return &Router{
		App:              fiber,
		Cfg:              cfg,
		AuthRouter:       authRouter,
		AttendanceRouter: attendanceRouter,
		EmployeeRouter:   employeeRouter,
		DepartmentRouter: departmentRouter,
		LocationRouter:   locationRouter,
		PositionRouter:   positionRouter,
		ReportRouter:     reportRouter,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Swagger Documentation
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	// Register routes of modules
	r.AuthRouter.RegisterAuthRoutes()

	r.App.Use(middleware.JWTMiddleware)
	r.AttendanceRouter.RegisterAttendanceRoutes()
	r.EmployeeRouter.RegisterEmployeeRoutes()
	r.DepartmentRouter.RegisterDepartmentRoutes()
	r.LocationRouter.RegisterLocationRoutes()
	r.PositionRouter.RegisterPositionRoutes()
	r.ReportRouter.RegisterReportRoutes()
}
