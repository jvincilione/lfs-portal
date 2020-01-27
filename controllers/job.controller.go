package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"lfs-portal/models"
	"lfs-portal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

type (
	JobController interface {
		GetJobById(c *gin.Context)
		GetAllJobs(c *gin.Context)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	job, err := controller.svc.GetJobById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resources": jobs})
}

func (controller jobController) CreateJob(c *gin.Context) {
	var job models.Job
	err := c.BindJSON(&job)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	v := validator.New()
	err = v.Struct(job)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	newJob, err := controller.svc.CreateJob(job)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusCreated, newJob)
}

func (controller jobController) UpdateJob(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	job, err := controller.svc.GetJobById(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if job.ID == 0 {
		logrus.Error(job.ID)
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	err = c.BindJSON(job)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	v := validator.New()
	err = v.Struct(job)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	updatedJob, err := controller.svc.UpdateJob(*job)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, updatedJob)
}

func (controller jobController) DeleteJob(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	err = controller.svc.DeleteJob(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
