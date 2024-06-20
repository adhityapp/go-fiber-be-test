package controller

import (
	"strconv"

	"github.com/bangadam/go-fiber-starter/app/module/employee/request"
	"github.com/bangadam/go-fiber-starter/app/module/employee/service"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
	"github.com/bangadam/go-fiber-starter/utils/response"
	"github.com/gofiber/fiber/v2"
)

type employeeController struct {
	employeeService service.EmployeeService
}

type EmployeeController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &employeeController{
		employeeService: employeeService,
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
func (_i *employeeController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.EmployeesRequest
	req.Pagination = paginate

	articles, paging, err := _i.employeeService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Employee list successfully retrieved"},
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
func (_i *employeeController) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	articles, err := _i.employeeService.Show(id)
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
func (_i *employeeController) Store(c *fiber.Ctx) error {

	req := new(request.EmployeeRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	req.CreatedBy = c.Locals("user").(string)

	err := _i.employeeService.Store(*req)
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
func (_i *employeeController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	req := new(request.EmployeeRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	req.UpdatedBy = c.Locals("user").(string)

	err = _i.employeeService.Update(id, *req)
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
func (_i *employeeController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	err = _i.employeeService.Destroy(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"successfully deleted"},
	})
}
