package repository

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/module/position/request"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type positionRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . ArticleRepository
type PositionRepository interface {
	Get(req request.PositionsRequest) (articles []*schema.Position, paging paginator.Pagination, err error)
	FindOne(id uint64) (article *schema.Position, err error)
	Create(article *schema.Position) (err error)
	Update(id uint64, article *schema.Position) (err error)
	Delete(id uint64) (err error)
}

func NewPositionRepository(db *database.Database) PositionRepository {
	return &positionRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *positionRepository) Get(req request.PositionsRequest) (articles []*schema.Position, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.Position{})
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

func (_i *positionRepository) FindOne(id uint64) (article *schema.Position, err error) {
	if err := _i.DB.DB.First(&article, id).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (_i *positionRepository) Create(article *schema.Position) (err error) {
	return _i.DB.DB.Create(article).Error
}

func (_i *positionRepository) Update(id uint64, article *schema.Position) (err error) {
	return _i.DB.DB.Model(&schema.Position{}).
		Where(&schema.Position{PositionID: id}).
		Updates(article).Error
}

func (_i *positionRepository) Delete(id uint64) error {
	return _i.DB.DB.Delete(&schema.Position{}, id).Error
}
