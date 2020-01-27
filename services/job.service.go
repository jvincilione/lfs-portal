package services

import (
	"fmt"
	"lfs-portal/models"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	JobService interface {
		GetJobById(ID int) (*models.Job, error)
		GetAllJobs() ([]models.ListJob, error)
		CreateJob(job models.Job) (*models.Job, error)
		UpdateJob(job models.Job) (*models.Job, error)
		DeleteJob(ID int) error
	}

	jobService struct {
		model         models.JobModel
	}
)

func NewJobService(model models.JobModel) JobService {
	return jobService{model}
}

func (svc jobService) GetJobById(ID int) (*models.Job, error) {
	job, err := svc.model.GetJobById(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetJobById] Error getting job, %v", err))
		return nil, err
	}
	return job, nil
}

func (svc jobService) GetAllJobs() ([]models.ListJob, error) {
	jobs, err := svc.model.GetAllJobs()
	if err != nil {
		logrus.Error(fmt.Sprintf("[GetAllJobs] Error getting all jobs, %v", err))
		return nil, err
	}
	return jobs, nil
}

func (svc jobService) CreateJob(job models.Job) (*models.Job, error) {
	var scheduledDate time.Time
	var err error
	if job.ScheduledDate != "" {
		scheduledDate, err = time.Parse("2006-01-02T15:04:05Z07:00", job.ScheduledDate)
	}

	if err != nil {
		logrus.Error(fmt.Sprintf("[CreateJob] Error converting schedule date for job, %v", err))
		return nil, err
	}

	if job.ScheduledDate != "" {
		job.ScheduledDate = scheduledDate.Format("2006-01-02 15:04")
	}

	newJob, err := svc.model.CreateJob(job)
	if err != nil {
		logrus.Error(fmt.Sprintf("[CreateJob] Error creating job, %v", err))
		return nil, err
	}
	return newJob, nil
}

func (svc jobService) UpdateJob(job models.Job) (*models.Job, error) {
	var scheduledDate time.Time
	var err error
	if job.ScheduledDate != "" {
		scheduledDate, err = time.Parse("2006-01-02T15:04:05Z07:00", job.ScheduledDate)
	}

	if err != nil {
		logrus.Error(fmt.Sprintf("[UpdateJob] Error converting schedule date for job, %v", err))
		return nil, err
	}

	if job.ScheduledDate != "" {
		job.ScheduledDate = scheduledDate.Format("2006-01-02 15:04")
	}

	updatedJob, err := svc.model.UpdateJob(job)
	if err != nil {
		logrus.Error(fmt.Sprintf("[UpdateJob] Error updating job, %v", err))
		return nil, err
	}
	return updatedJob, nil
}

func (svc jobService) DeleteJob(ID int) error {
	var jobID uint
	jobID = uint(ID)
	err := svc.model.DeleteJob(jobID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[DeleteJob] Error deleting job, %v", err))
		return err
	}
	return nil
}
