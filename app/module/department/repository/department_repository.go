package repository

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/module/department/request"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type departmentRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . ArticleRepository
type DepartmentRepository interface {
	Get(req request.DepartmentsRequest) (articles []*schema.Department, paging paginator.Pagination, err error)
	FindOne(id uint64) (article *schema.Department, err error)
	Create(article *schema.Department) (err error)
	Update(id uint64, article *schema.Department) (err error)
	Delete(id uint64) (err error)
}

func NewDepartmentRepository(db *database.Database) DepartmentRepository {
	return &departmentRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *departmentRepository) Get(req request.DepartmentsRequest) (articles []*schema.Department, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.Department{})
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

func (_i *departmentRepository) FindOne(id uint64) (article *schema.Department, err error) {
	if err := _i.DB.DB.First(&article, id).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (_i *departmentRepository) Create(article *schema.Department) (err error) {
	return _i.DB.DB.Create(article).Error
}

func (_i *departmentRepository) Update(id uint64, article *schema.Department) (err error) {
	return _i.DB.DB.Model(&schema.Department{}).
		Where(&schema.Department{DepartmentID: id}).
		Updates(article).Error
}

func (_i *departmentRepository) Delete(id uint64) error {
	return _i.DB.DB.Delete(&schema.Department{}, id).Error
}
