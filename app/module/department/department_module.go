package department

import (
	"github.com/bangadam/go-fiber-starter/app/module/department/controller"
	"github.com/bangadam/go-fiber-starter/app/module/department/repository"
	"github.com/bangadam/go-fiber-starter/app/module/department/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of ArticleRouter
type DepartmentRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of article module
var NewDepartmentModule = fx.Options(
	// register repository of article module
	fx.Provide(repository.NewDepartmentRepository),

	// register service of article module
	fx.Provide(service.NewDepartmentService),

	// register controller of article module
	fx.Provide(controller.NewController),

	// register router of article module
	fx.Provide(NewDepartmentRouter),
)

// init ArticleRouter
func NewDepartmentRouter(fiber *fiber.App, controller *controller.Controller) *DepartmentRouter {
	return &DepartmentRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of article module
func (_i *DepartmentRouter) RegisterDepartmentRoutes() {
	// define controllers
	departmentController := _i.Controller.Department

	// define routes
	_i.App.Route("/department", func(router fiber.Router) {
		router.Get("/", departmentController.Index)
		router.Get("/:id", departmentController.Show)
		router.Post("/", departmentController.Store)
		router.Put("/:id", departmentController.Update)
		router.Delete("/:id", departmentController.Delete)
	})
}
