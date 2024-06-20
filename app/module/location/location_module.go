package location

import (
	"github.com/bangadam/go-fiber-starter/app/module/location/controller"
	"github.com/bangadam/go-fiber-starter/app/module/location/repository"
	"github.com/bangadam/go-fiber-starter/app/module/location/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of ArticleRouter
type LocationRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of article module
var NewLocationModule = fx.Options(
	// register repository of article module
	fx.Provide(repository.NewLocationRepository),

	// register service of article module
	fx.Provide(service.NewLocationService),

	// register controller of article module
	fx.Provide(controller.NewController),

	// register router of article module
	fx.Provide(NewLocationRouter),
)

// init ArticleRouter
func NewLocationRouter(fiber *fiber.App, controller *controller.Controller) *LocationRouter {
	return &LocationRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of article module
func (_i *LocationRouter) RegisterLocationRoutes() {
	// define controllers
	locationController := _i.Controller.Location

	// define routes
	_i.App.Route("/location", func(router fiber.Router) {
		router.Get("/", locationController.Index)
		router.Get("/:id", locationController.Show)
		router.Post("/", locationController.Store)
		router.Put("/:id", locationController.Update)
		router.Delete("/:id", locationController.Delete)
	})
}
