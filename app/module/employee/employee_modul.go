package employee

import (
	"github.com/bangadam/go-fiber-starter/app/module/employee/controller"
	"github.com/bangadam/go-fiber-starter/app/module/employee/repository"
	"github.com/bangadam/go-fiber-starter/app/module/employee/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of ArticleRouter
type EmployeeRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of article module
var NewEmployeeModule = fx.Options(
	// register repository of article module
	fx.Provide(repository.NewEmployeeRepository),

	// register service of article module
	fx.Provide(service.NewEmployeeService),

	// register controller of article module
	fx.Provide(controller.NewController),

	// register router of article module
	fx.Provide(NewEmployeeRouter),
)

// init ArticleRouter
func NewEmployeeRouter(fiber *fiber.App, controller *controller.Controller) *EmployeeRouter {
	return &EmployeeRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of article module
func (_i *EmployeeRouter) RegisterEmployeeRoutes() {
	// define controllers
	employeeController := _i.Controller.Employee

	// define routes
	_i.App.Route("/employee", func(router fiber.Router) {
		router.Get("/", employeeController.Index)
		router.Get("/:id", employeeController.Show)
		router.Post("/", employeeController.Store)
		router.Put("/:id", employeeController.Update)
		router.Delete("/:id", employeeController.Delete)
	})
}
