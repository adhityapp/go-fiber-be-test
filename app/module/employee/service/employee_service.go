package service

import (
	"fmt"
	"time"

	"github.com/bangadam/go-fiber-starter/app/module/employee/repository"
	"github.com/bangadam/go-fiber-starter/app/module/employee/request"
	"github.com/bangadam/go-fiber-starter/app/module/employee/response"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

// ArticleService
type employeeService struct {
	Repo repository.EmployeeRepository
}

// define interface of IArticleService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . ArticleService
type EmployeeService interface {
	All(req request.EmployeesRequest) (articles []*response.Employee, paging paginator.Pagination, err error)
	Show(id uint64) (article *response.Employee, err error)
	Store(req request.EmployeeRequest) (err error)
	Update(id uint64, req request.EmployeeRequest) (err error)
	Destroy(id uint64) error
}

// init ArticleService
func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		Repo: repo,
	}
}

// implement interface of ArticleService
func (_i *employeeService) All(req request.EmployeesRequest) (articles []*response.Employee, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.Get(req)
	if err != nil {
		return
	}

	for _, result := range results {
		articles = append(articles, response.FromDomain(result))
	}

	return
}

func (_i *employeeService) Show(id uint64) (article *response.Employee, err error) {
	result, err := _i.Repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *employeeService) Store(req request.EmployeeRequest) (err error) {
	// Increment to generate the next number
	val, err := _i.Repo.GenerateEmployeeCode()
	if err != nil {
		return err
	}
	currentYearMonth := time.Now().Format("0601")

	nextNumber := val.EmployeeID + 1
	employeeCode := fmt.Sprintf("%s%04d", currentYearMonth, nextNumber)

	req.EmployeeCode = employeeCode

	return _i.Repo.Create(req.ToDomain())
}

func (_i *employeeService) Update(id uint64, req request.EmployeeRequest) (err error) {
	return _i.Repo.Update(id, req.ToDomain())
}

func (_i *employeeService) Destroy(id uint64) error {
	return _i.Repo.Delete(id)
}
