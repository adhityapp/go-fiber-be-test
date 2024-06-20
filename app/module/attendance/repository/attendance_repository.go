package repository

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/module/attendance/request"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type attendanceRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . ArticleRepository
type AttendanceRepository interface {
	Get(req request.AttendancesRequest) (articles []*schema.Attendance, paging paginator.Pagination, err error)
	FindOne(id uint64) (article *schema.Attendance, err error)
	Create(article *schema.Attendance) (err error)
	Update(id uint64, article *schema.Attendance) (err error)
	Delete(id uint64) (err error)
}

func NewAttendanceRepository(db *database.Database) AttendanceRepository {
	return &attendanceRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *attendanceRepository) Get(req request.AttendancesRequest) (articles []*schema.Attendance, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.Attendance{})
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

func (_i *attendanceRepository) FindOne(id uint64) (article *schema.Attendance, err error) {
	if err := _i.DB.DB.First(&article, id).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (_i *attendanceRepository) Create(article *schema.Attendance) (err error) {
	return _i.DB.DB.Create(article).Error
}

func (_i *attendanceRepository) Update(id uint64, article *schema.Attendance) (err error) {
	return _i.DB.DB.Model(&schema.Attendance{}).
		Where(&schema.Attendance{AttendanceID: id}).
		Updates(article).Error
}

func (_i *attendanceRepository) Delete(id uint64) error {
	return _i.DB.DB.Delete(&schema.Attendance{}, id).Error
}
