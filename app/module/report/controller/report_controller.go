package controller

import (
	"github.com/bangadam/go-fiber-starter/app/module/report/request"
	"github.com/bangadam/go-fiber-starter/app/module/report/service"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
	"github.com/bangadam/go-fiber-starter/utils/response"
	"github.com/gofiber/fiber/v2"
)

type reportController struct {
	reportService service.ReportService
}

type ReportController interface {
	Index(c *fiber.Ctx) error
	// Show(c *fiber.Ctx) error
	// Store(c *fiber.Ctx) error
	// Update(c *fiber.Ctx) error
	// Delete(c *fiber.Ctx) error
}

func NewReportController(reportService service.ReportService) ReportController {
	return &reportController{
		reportService: reportService,
	}
}

// get all articles
// @Summary      Get all articles
// @Description  API for getting all articles
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles [get]
func (_i *reportController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.ReportsRequest
	req.Pagination = paginate

	articles, paging, err := _i.reportService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Report list successfully retrieved"},
		Data:     articles,
		Meta:     paging,
	})
}

// get one article
// @Summary      Get one article
// @Description  API for getting one article
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Article ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles/:id [get]
// func (_i *reportController) Show(c *fiber.Ctx) error {
// 	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	articles, err := _i.reportService.Show(id)
// 	if err != nil {
// 		return err
// 	}

// 	return response.Resp(c, response.Response{
// 		Messages: response.Messages{"successfully retrieved"},
// 		Data:     articles,
// 	})
// }

// // create article
// // @Summary      Create article
// // @Description  API for create article
// // @Tags         Task
// // @Security     Bearer
// // @Body 	     request.ArticleRequest
// // @Success      200  {object}  response.Response
// // @Failure      401  {object}  response.Response
// // @Failure      404  {object}  response.Response
// // @Failure      422  {object}  response.Response
// // @Failure      500  {object}  response.Response
// // @Router       /articles [post]
// func (_i *reportController) Store(c *fiber.Ctx) error {
// 	req := new(request.ReportRequest)
// 	if err := response.ParseAndValidate(c, req); err != nil {
// 		return err
// 	}

// 	req.AbsenIN = time.Now()
// 	req.CreatedBy = c.Locals("user").(string)

// 	err := _i.reportService.Store(*req)
// 	if err != nil {
// 		return err
// 	}

// 	return response.Resp(c, response.Response{
// 		Messages: response.Messages{"successfully created"},
// 	})
// }

// // update article
// // @Summary      update article
// // @Description  API for update article
// // @Tags         Task
// // @Security     Bearer
// // @Body 	     request.ArticleRequest
// // @Param        id path int true "Article ID"
// // @Success      200  {object}  response.Response
// // @Failure      401  {object}  response.Response
// // @Failure      404  {object}  response.Response
// // @Failure      422  {object}  response.Response
// // @Failure      500  {object}  response.Response
// // @Router       /articles/:id [put]
// func (_i *reportController) Update(c *fiber.Ctx) error {
// 	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	req := new(request.ReportRequest)
// 	if err := response.ParseAndValidate(c, req); err != nil {
// 		return err
// 	}

// 	req.AbsenOUT = time.Now()
// 	req.UpdatedBy = c.Locals("user").(string)

// 	err = _i.reportService.Update(id, *req)
// 	if err != nil {
// 		return err
// 	}

// 	return response.Resp(c, response.Response{
// 		Messages: response.Messages{"successfully updated"},
// 	})
// }

// // delete article
// // @Summary      delete article
// // @Description  API for delete article
// // @Tags         Task
// // @Security     Bearer
// // @Param        id path int true "Article ID"
// // @Success      200  {object}  response.Response
// // @Failure      401  {object}  response.Response
// // @Failure      404  {object}  response.Response
// // @Failure      422  {object}  response.Response
// // @Failure      500  {object}  response.Response
// // @Router       /articles/:id [delete]
// func (_i *reportController) Delete(c *fiber.Ctx) error {
// 	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	err = _i.reportService.Destroy(id)
// 	if err != nil {
// 		return err
// 	}

// 	return response.Resp(c, response.Response{
// 		Messages: response.Messages{"successfully deleted"},
// 	})
// }
