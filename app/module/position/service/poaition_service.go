package service

import (
	"github.com/bangadam/go-fiber-starter/app/module/position/repository"
	"github.com/bangadam/go-fiber-starter/app/module/position/request"
	"github.com/bangadam/go-fiber-starter/app/module/position/response"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

// ArticleService
type positionService struct {
	Repo repository.PositionRepository
}

// define interface of IArticleService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . ArticleService
type PositionService interface {
	All(req request.PositionsRequest) (articles []*response.Position, paging paginator.Pagination, err error)
	Show(id uint64) (article *response.Position, err error)
	Store(req request.PositionRequest) (err error)
	Update(id uint64, req request.PositionRequest) (err error)
	Destroy(id uint64) error
}

// init ArticleService
func NewPositionService(repo repository.PositionRepository) PositionService {
	return &positionService{
		Repo: repo,
	}
}

// implement interface of ArticleService
func (_i *positionService) All(req request.PositionsRequest) (articles []*response.Position, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.Get(req)
	if err != nil {
		return
	}

	for _, result := range results {
		articles = append(articles, response.FromDomain(result))
	}

	return
}

func (_i *positionService) Show(id uint64) (article *response.Position, err error) {
	result, err := _i.Repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *positionService) Store(req request.PositionRequest) (err error) {
	return _i.Repo.Create(req.ToDomain())
}

func (_i *positionService) Update(id uint64, req request.PositionRequest) (err error) {
	return _i.Repo.Update(id, req.ToDomain())
}

func (_i *positionService) Destroy(id uint64) error {
	return _i.Repo.Delete(id)
}
