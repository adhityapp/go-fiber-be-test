package service

import (
	"github.com/bangadam/go-fiber-starter/app/module/report/repository"
	"github.com/bangadam/go-fiber-starter/app/module/report/request"
	"github.com/bangadam/go-fiber-starter/app/module/report/response"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

// ArticleService
type reportService struct {
	Repo repository.ReportRepository
}

// define interface of IArticleService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . ArticleService
type ReportService interface {
	All(req request.ReportsRequest) (articles []*response.Report, paging paginator.Pagination, err error)
	// Show(id uint64) (article *response.Report, err error)
	// Store(req request.ReportRequest) (err error)
	// Update(id uint64, req request.ReportRequest) (err error)
	// Destroy(id uint64) error
}

// init ArticleService
func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{
		Repo: repo,
	}
}

// implement interface of ArticleService
func (_i *reportService) All(req request.ReportsRequest) (articles []*response.Report, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.Get(req)
	if err != nil {
		return
	}

	for _, result := range results {
		articles = append(articles, response.FromDomain(result))
	}

	return
}

// func (_i *reportService) Show(id uint64) (article *response.Report, err error) {
// 	result, err := _i.Repo.FindOne(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return response.FromDomain(result), nil
// }

// func (_i *reportService) Store(req request.ReportRequest) (err error) {
// 	return _i.Repo.Create(req.ToDomain())
// }

// func (_i *reportService) Update(id uint64, req request.ReportRequest) (err error) {
// 	return _i.Repo.Update(id, req.ToDomain())
// }

// func (_i *reportService) Destroy(id uint64) error {
// 	return _i.Repo.Delete(id)
// }
