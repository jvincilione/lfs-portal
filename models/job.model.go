package models

import (
	"lfs-portal/enums"
	"time"

	"github.com/jinzhu/gorm"
)

type (
	Job struct {
		gorm.Model
		FullName        string          `json:"fullName,omitempty" validate:"required" gorm:"type:varchar(100);not null"`
		Address         string          `json:"address,omitempty" validate:"required" gorm:"type:varchar(255);not null"`
		Address2        string          `json:"address2,omitempty" gorm:"type:varchar(255)"`
		City            string          `json:"city,omitempty" gorm:"type:varchar(255)"`
		State           string          `json:"state,omitempty" gorm:"type:varchar(15)"`
		PostalCode      string          `json:"postalCode,omitempty" validate:"required" gorm:"type:varchar(10);not null"`
		PhoneNumber     string          `json:"phoneNumber,omitempty" validate:"required" gorm:"type:varchar(11);not null"`
		OrderNumber     string          `json:"orderNumber,omitempty" gorm:"type:varchar(255)"`
		Instructions    string          `json:"instructions,omitempty" gorm:"type:text"`
		ScheduledDate   *time.Time      `json:"scheduledDate,omitempty" gorm:"type:datetime"`
		CompletededDate *time.Time      `json:"completedDate,omitempty" gorm:"type:datetime"`
		PaymentDate     *time.Time      `json:"paymentDate,omitempty" gorm:"type:datetime"`
		Status          enums.JobStatus `json:"status,omitempty" gorm:"type:int;not null"`
		PartsCost       float64         `json:"partsCost,omitempty" gorm:"type:float"`
		LaborCost       float64         `json:"laborCost,omitempty" gorm:"type:float"`
		Notes           string          `json:"notes,omitempty" validate:"-" gorm:"type:text"`
		CompanyID       uint            `json:"companyId,omitempty" validate:"required" gorm:"type:int(10)unsigned"`
		Company         Company         `json:"company" validate:"-" gorm:"association_autoupdate:false "`
	}

	ListJob struct {
		ID            uint
		CreatedAt     time.Time
		UpdatedAt     time.Time
		FullName      string          `json:"fullName"`
		City          string          `json:"city"`
		State         string          `json:"state"`
		PostalCode    string          `json:"postalCode"`
		PhoneNumber   string          `json:"phoneNumber"`
		ScheduledDate string          `json:"scheduledDate"`
		Status        enums.JobStatus `json:"status"`
		Company       Company         `json:"company"`
	}

	JobModel interface {
		GetJobById(ID int) (*Job, error)
		GetAllJobs() ([]ListJob, error)
		GetCompanyJobs(companyID uint) ([]ListJob, error)
		CreateJob(job Job) (*Job, error)
		UpdateJob(job Job) (*Job, error)
		DeleteJob(ID uint) error
	}

	jobModel struct {
		db *gorm.DB
	}
)

func NewJob(db *gorm.DB) JobModel {
	return jobModel{db}
}

func (model jobModel) GetJobById(ID int) (*Job, error) {
	var job Job
	err := model.db.Preload("Company").First(&job, ID).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (model jobModel) GetAllJobs() ([]ListJob, error) {
	var jobList []ListJob
	err := model.db.Table("jobs").Where("status != ? && status != ?", 4, 5).Find(&jobList).Error

	if err != nil {
		return nil, err
	}
	return jobList, nil
}

func (model jobModel) GetCompanyJobs(companyID uint) ([]ListJob, error) {
	var jobList []ListJob
	err := model.db.Table("jobs").Where("company_id = ?", companyID).Find(&jobList).Error

	if err != nil {
		return nil, err
	}
	return jobList, nil
}

func (model jobModel) CreateJob(job Job) (*Job, error) {
	newJob := Job{}

	err := model.db.Create(&job).Scan(&newJob).Error

	if err != nil {
		return nil, err
	}

	model.db.Model(&newJob).Related(&newJob.Company)

	return &newJob, nil
}

func (model jobModel) UpdateJob(job Job) (*Job, error) {
	updatedJob := Job{}

	err := model.db.Save(&job).Scan(&updatedJob).Error

	if err != nil {
		return nil, err
	}
	model.db.Model(&updatedJob).Related(&updatedJob.Company)
	return &updatedJob, nil
}

func (model jobModel) DeleteJob(ID uint) error {
	err := model.db.Unscoped().Delete(Job{Model: gorm.Model{ID: ID}}).Error

	if err != nil {
		return err
	}
	return nil
}
