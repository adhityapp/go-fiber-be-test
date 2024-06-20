package repository

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/module/location/request"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type locationRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . ArticleRepository
type LocationRepository interface {
	Get(req request.LocationsRequest) (articles []*schema.Location, paging paginator.Pagination, err error)
	FindOne(id uint64) (article *schema.Location, err error)
	Create(article *schema.Location) (err error)
	Update(id uint64, article *schema.Location) (err error)
	Delete(id uint64) (err error)
}

func NewLocationRepository(db *database.Database) LocationRepository {
	return &locationRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *locationRepository) Get(req request.LocationsRequest) (articles []*schema.Location, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.Location{})
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

func (_i *locationRepository) FindOne(id uint64) (article *schema.Location, err error) {
	if err := _i.DB.DB.First(&article, id).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (_i *locationRepository) Create(article *schema.Location) (err error) {
	return _i.DB.DB.Create(article).Error
}

func (_i *locationRepository) Update(id uint64, article *schema.Location) (err error) {
	return _i.DB.DB.Model(&schema.Location{}).
		Where(&schema.Location{LocationID: id}).
		Updates(article).Error
}

func (_i *locationRepository) Delete(id uint64) error {
	return _i.DB.DB.Delete(&schema.Location{}, id).Error
}
