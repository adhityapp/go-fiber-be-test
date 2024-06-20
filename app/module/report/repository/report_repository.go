package repository

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/module/report/request"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type reportRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . ArticleRepository
type ReportRepository interface {
	Get(req request.ReportsRequest) (articles []*schema.Report, paging paginator.Pagination, err error)
	// FindOne(id uint64) (article *schema.Report, err error)
	// Create(article *schema.Report) (err error)
	// Update(id uint64, article *schema.Report) (err error)
	// Delete(id uint64) (err error)
}

func NewReportRepository(db *database.Database) ReportRepository {
	return &reportRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *reportRepository) Get(req request.ReportsRequest) (articles []*schema.Report, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Table("attendances").
		Select("attendances.attendance_id, attendances.absen_in as date,employees.employee_code, employees.employee_name, departments.department_name, positions.position_name, locations.location_name, attendances.absen_in, attendances.absen_out").
		Joins("left join employees on employees.employee_id = attendances.employee_id").
		Joins("left join locations on locations.location_id = attendances.location_id").
		Joins("left join departments on departments.department_id = employees.department_id").
		Joins("left join positions on positions.position_id = employees.position_id")

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

// func (_i *reportRepository) FindOne(id uint64) (article *schema.Report, err error) {
// 	if err := _i.DB.DB.First(&article, id).Error; err != nil {
// 		return nil, err
// 	}

// 	return article, nil
// }

// func (_i *reportRepository) Create(article *schema.Report) (err error) {
// 	return _i.DB.DB.Create(article).Error
// }

// func (_i *reportRepository) Update(id uint64, article *schema.Report) (err error) {
// 	return _i.DB.DB.Model(&schema.Report{}).
// 		Where(&schema.Report{ReportID: id}).
// 		Updates(article).Error
// }

// func (_i *reportRepository) Delete(id uint64) error {
// 	return _i.DB.DB.Delete(&schema.Report{}, id).Error
// }
