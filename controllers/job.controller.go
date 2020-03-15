package controllers

import (
	"lfs-portal/models"
	"lfs-portal/services"
	"lfs-portal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

type (
	JobController interface {
		GetJobById(c *gin.Context)
		GetAllJobs(c *gin.Context)
		GetCompanyJobs(c *gin.Context)
		CreateJob(c *gin.Context)
		UpdateJob(c *gin.Context)
		DeleteJob(c *gin.Context)
	}

	jobController struct {
		svc services.JobService
	}
)

func NewJobController(svc services.JobService) JobController {
	return jobController{svc}
}

func (controller jobController) GetJobById(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}
	job, err := controller.svc.GetJobById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	if job.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, job)
}

func (controller jobController) GetAllJobs(c *gin.Context) {
	jobs, err := controller.svc.GetAllJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"resources": jobs})
}

func (controller jobController) GetCompanyJobs(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}
	jobs, err := controller.svc.GetCompanyJobs(uint(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"resources": jobs})
}

func (controller jobController) CreateJob(c *gin.Context) {
	var job models.Job
	err := c.BindJSON(&job)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	v := validator.New()
	err = v.Struct(job)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	newJob, err := controller.svc.CreateJob(job)
	if err != nil {
		c.JSON(http.StatusConflict, utils.GetConflictError(err))
		return
	}
	c.JSON(http.StatusCreated, newJob)
}

func (controller jobController) UpdateJob(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}
	job, err := controller.svc.GetJobById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	if job.ID == 0 {
		logrus.Error(job.ID)
		c.JSON(http.StatusNotFound, utils.GetNotFoundError())
		return
	}

	err = c.BindJSON(job)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	v := validator.New()
	err = v.Struct(job)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}

	updatedJob, err := controller.svc.UpdateJob(*job)
	if err != nil {
		c.JSON(http.StatusConflict, utils.GetConflictError(err))
		return
	}
	c.JSON(http.StatusOK, updatedJob)
}

func (controller jobController) DeleteJob(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GetBadRequestError(err))
		return
	}
	err = controller.svc.DeleteJob(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GetGenericError(err))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
