package service

import (
	"github.com/bangadam/go-fiber-starter/app/module/department/repository"
	"github.com/bangadam/go-fiber-starter/app/module/department/request"
	"github.com/bangadam/go-fiber-starter/app/module/department/response"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

// ArticleService
type departmentService struct {
	Repo repository.DepartmentRepository
}

// define interface of IArticleService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . ArticleService
type DepartmentService interface {
	All(req request.DepartmentsRequest) (articles []*response.Department, paging paginator.Pagination, err error)
	Show(id uint64) (article *response.Department, err error)
	Store(req request.DepartmentRequest) (err error)
	Update(id uint64, req request.DepartmentRequest) (err error)
	Destroy(id uint64) error
}

// init ArticleService
func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
	return &departmentService{
		Repo: repo,
	}
}

// implement interface of ArticleService
func (_i *departmentService) All(req request.DepartmentsRequest) (articles []*response.Department, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.Get(req)
	if err != nil {
		return
	}

	for _, result := range results {
		articles = append(articles, response.FromDomain(result))
	}

	return
}

func (_i *departmentService) Show(id uint64) (article *response.Department, err error) {
	result, err := _i.Repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *departmentService) Store(req request.DepartmentRequest) (err error) {
	return _i.Repo.Create(req.ToDomain())
}

func (_i *departmentService) Update(id uint64, req request.DepartmentRequest) (err error) {
	return _i.Repo.Update(id, req.ToDomain())
}

func (_i *departmentService) Destroy(id uint64) error {
	return _i.Repo.Delete(id)
}
