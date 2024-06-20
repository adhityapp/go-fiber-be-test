package report

import (
	"github.com/bangadam/go-fiber-starter/app/module/report/controller"
	"github.com/bangadam/go-fiber-starter/app/module/report/repository"
	"github.com/bangadam/go-fiber-starter/app/module/report/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of ArticleRouter
type ReportRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of article module
var NewReportModule = fx.Options(
	// register repository of article module
	fx.Provide(repository.NewReportRepository),

	// register service of article module
	fx.Provide(service.NewReportService),

	// register controller of article module
	fx.Provide(controller.NewController),

	// register router of article module
	fx.Provide(NewReportRouter),
)

// init ArticleRouter
func NewReportRouter(fiber *fiber.App, controller *controller.Controller) *ReportRouter {
	return &ReportRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of article module
func (_i *ReportRouter) RegisterReportRoutes() {
	// define controllers
	reportController := _i.Controller.Report

	// define routes
	_i.App.Route("/report", func(router fiber.Router) {
		router.Get("/", reportController.Index)
		// router.Get("/:id", reportController.Show)
		// router.Post("/", reportController.Store)
		// router.Put("/:id", reportController.Update)
		// router.Delete("/:id", reportController.Delete)
	})
}
