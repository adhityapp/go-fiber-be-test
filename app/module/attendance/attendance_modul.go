package attendance

import (
	"github.com/bangadam/go-fiber-starter/app/module/attendance/controller"
	"github.com/bangadam/go-fiber-starter/app/module/attendance/repository"
	"github.com/bangadam/go-fiber-starter/app/module/attendance/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of ArticleRouter
type AttendanceRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of article module
var NewAttendanceModule = fx.Options(
	// register repository of article module
	fx.Provide(repository.NewAttendanceRepository),

	// register service of article module
	fx.Provide(service.NewAttendanceService),

	// register controller of article module
	fx.Provide(controller.NewController),

	// register router of article module
	fx.Provide(NewAttendanceRouter),
)

// init ArticleRouter
func NewAttendanceRouter(fiber *fiber.App, controller *controller.Controller) *AttendanceRouter {
	return &AttendanceRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of article module
func (_i *AttendanceRouter) RegisterAttendanceRoutes() {
	// define controllers
	attendanceController := _i.Controller.Attendance

	// define routes
	_i.App.Route("/attendance", func(router fiber.Router) {
		router.Get("/", attendanceController.Index)
		router.Get("/:id", attendanceController.Show)
		router.Post("/", attendanceController.Store)
		router.Put("/:id", attendanceController.Update)
		router.Delete("/:id", attendanceController.Delete)
	})
}
