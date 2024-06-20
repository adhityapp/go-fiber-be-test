package controller

import (
	"strconv"
	"time"

	"github.com/bangadam/go-fiber-starter/app/module/attendance/request"
	"github.com/bangadam/go-fiber-starter/app/module/attendance/service"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
	"github.com/bangadam/go-fiber-starter/utils/response"
	"github.com/gofiber/fiber/v2"
)

type attendanceController struct {
	attendanceService service.AttendanceService
}

type AttendanceController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewAttendanceController(attendanceService service.AttendanceService) AttendanceController {
	return &attendanceController{
		attendanceService: attendanceService,
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
func (_i *attendanceController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.AttendancesRequest
	req.Pagination = paginate

	articles, paging, err := _i.attendanceService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Attendance list successfully retrieved"},
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
func (_i *attendanceController) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	articles, err := _i.attendanceService.Show(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"successfully retrieved"},
		Data:     articles,
	})
}

// create article
// @Summary      Create article
// @Description  API for create article
// @Tags         Task
// @Security     Bearer
// @Body 	     request.ArticleRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles [post]
func (_i *attendanceController) Store(c *fiber.Ctx) error {
	req := new(request.AttendanceRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	req.AbsenIN = time.Now()
	req.CreatedBy = c.Locals("user").(string)

	err := _i.attendanceService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"successfully created"},
	})
}

// update article
// @Summary      update article
// @Description  API for update article
// @Tags         Task
// @Security     Bearer
// @Body 	     request.ArticleRequest
// @Param        id path int true "Article ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles/:id [put]
func (_i *attendanceController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	req := new(request.AttendanceRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	req.AbsenOUT = time.Now()
	req.UpdatedBy = c.Locals("user").(string)

	err = _i.attendanceService.Update(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"successfully updated"},
	})
}

// delete article
// @Summary      delete article
// @Description  API for delete article
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Article ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles/:id [delete]
func (_i *attendanceController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	err = _i.attendanceService.Destroy(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"successfully deleted"},
	})
}
