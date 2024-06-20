package position

import (
	"github.com/bangadam/go-fiber-starter/app/module/position/controller"
	"github.com/bangadam/go-fiber-starter/app/module/position/repository"
	"github.com/bangadam/go-fiber-starter/app/module/position/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of ArticleRouter
type PositionRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of article module
var NewPositionModule = fx.Options(
	// register repository of article module
	fx.Provide(repository.NewPositionRepository),

	// register service of article module
	fx.Provide(service.NewPositionService),

	// register controller of article module
	fx.Provide(controller.NewController),

	// register router of article module
	fx.Provide(NewPositionRouter),
)

// init ArticleRouter
func NewPositionRouter(fiber *fiber.App, controller *controller.Controller) *PositionRouter {
	return &PositionRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of article module
func (_i *PositionRouter) RegisterPositionRoutes() {
	// define controllers
	positionController := _i.Controller.Position

	// define routes
	_i.App.Route("/position", func(router fiber.Router) {
		router.Get("/", positionController.Index)
		router.Get("/:id", positionController.Show)
		router.Post("/", positionController.Store)
		router.Put("/:id", positionController.Update)
		router.Delete("/:id", positionController.Delete)
	})
}
