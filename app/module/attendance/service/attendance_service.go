package service

import (
	"github.com/bangadam/go-fiber-starter/app/module/attendance/repository"
	"github.com/bangadam/go-fiber-starter/app/module/attendance/request"
	"github.com/bangadam/go-fiber-starter/app/module/attendance/response"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

// ArticleService
type attendanceService struct {
	Repo repository.AttendanceRepository
}

// define interface of IArticleService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . ArticleService
type AttendanceService interface {
	All(req request.AttendancesRequest) (articles []*response.Attendance, paging paginator.Pagination, err error)
	Show(id uint64) (article *response.Attendance, err error)
	Store(req request.AttendanceRequest) (err error)
	Update(id uint64, req request.AttendanceRequest) (err error)
	Destroy(id uint64) error
}

// init ArticleService
func NewAttendanceService(repo repository.AttendanceRepository) AttendanceService {
	return &attendanceService{
		Repo: repo,
	}
}

// implement interface of ArticleService
func (_i *attendanceService) All(req request.AttendancesRequest) (articles []*response.Attendance, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.Get(req)
	if err != nil {
		return
	}

	for _, result := range results {
		articles = append(articles, response.FromDomain(result))
	}

	return
}

func (_i *attendanceService) Show(id uint64) (article *response.Attendance, err error) {
	result, err := _i.Repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *attendanceService) Store(req request.AttendanceRequest) (err error) {
	return _i.Repo.Create(req.ToDomain())
}

func (_i *attendanceService) Update(id uint64, req request.AttendanceRequest) (err error) {
	return _i.Repo.Update(id, req.ToDomain())
}

func (_i *attendanceService) Destroy(id uint64) error {
	return _i.Repo.Delete(id)
}
