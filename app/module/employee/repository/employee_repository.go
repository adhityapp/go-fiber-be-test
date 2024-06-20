package repository

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/module/employee/request"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type employeeRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . ArticleRepository
type EmployeeRepository interface {
	Get(req request.EmployeesRequest) (articles []*schema.Employee, paging paginator.Pagination, err error)
	FindOne(id uint64) (article *schema.Employee, err error)
	Create(article *schema.Employee) (err error)
	Update(id uint64, article *schema.Employee) (err error)
	Delete(id uint64) (err error)
	GenerateEmployeeCode() (article *schema.Employee, err error)
}

func NewEmployeeRepository(db *database.Database) EmployeeRepository {
	return &employeeRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *employeeRepository) Get(req request.EmployeesRequest) (articles []*schema.Employee, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.Employee{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&articles).Error
	if err != nil {
		return
	}

	paging = *req.Pagination

	return
}

func (_i *employeeRepository) FindOne(id uint64) (article *schema.Employee, err error) {
	if err := _i.DB.DB.First(&article, id).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (_i *employeeRepository) Create(article *schema.Employee) (err error) {
	return _i.DB.DB.Create(article).Error
}

func (_i *employeeRepository) Update(id uint64, article *schema.Employee) (err error) {
	return _i.DB.DB.Model(&schema.Employee{}).
		Where(&schema.Employee{EmployeeID: id}).
		Updates(article).Error
}

func (_i *employeeRepository) Delete(id uint64) error {
	return _i.DB.DB.Delete(&schema.Employee{}, id).Error
}

func (_i *employeeRepository) GenerateEmployeeCode() (article *schema.Employee, err error) {
	query := _i.DB.DB.Model(&schema.Employee{})
	query.Order("employee_id DESC")
	err = query.First(&article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}
