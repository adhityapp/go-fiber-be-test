package service

import (
	"github.com/bangadam/go-fiber-starter/app/module/location/repository"
	"github.com/bangadam/go-fiber-starter/app/module/location/request"
	"github.com/bangadam/go-fiber-starter/app/module/location/response"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

// ArticleService
type locationService struct {
	Repo repository.LocationRepository
}

// define interface of IArticleService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . ArticleService
type LocationService interface {
	All(req request.LocationsRequest) (articles []*response.Location, paging paginator.Pagination, err error)
	Show(id uint64) (article *response.Location, err error)
	Store(req request.LocationRequest) (err error)
	Update(id uint64, req request.LocationRequest) (err error)
	Destroy(id uint64) error
}

// init ArticleService
func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{
		Repo: repo,
	}
}

// implement interface of ArticleService
func (_i *locationService) All(req request.LocationsRequest) (articles []*response.Location, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.Get(req)
	if err != nil {
		return
	}

	for _, result := range results {
		articles = append(articles, response.FromDomain(result))
	}

	return
}

func (_i *locationService) Show(id uint64) (article *response.Location, err error) {
	result, err := _i.Repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *locationService) Store(req request.LocationRequest) (err error) {
	return _i.Repo.Create(req.ToDomain())
}

func (_i *locationService) Update(id uint64, req request.LocationRequest) (err error) {
	return _i.Repo.Update(id, req.ToDomain())
}

func (_i *locationService) Destroy(id uint64) error {
	return _i.Repo.Delete(id)
}
