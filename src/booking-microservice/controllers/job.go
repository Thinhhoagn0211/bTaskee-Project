package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/daos"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/models"
)

// Job manages Job CRUD
type Job struct {
	jobDAO daos.Jobs
}

// BookJob godoc
// @Summary Book a new job
// @Description Book a new job
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} models.Message
// @Router /Job/AddJob [post]
func (m *Job) BookJob(ctx *gin.Context) {

	var newJob daos.Jobs
	if err := ctx.ShouldBindJSON(&newJob.Job); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(newJob.Job)
	err := m.jobDAO.AddJobToDb(newJob)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Job booked successfully"})
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// RemoveJob godoc
// @Summary Remove a job from the list
// @Description Remove a job from the list
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} models.Message
// @Router /Job/RemoveJob [delete]
func (m *Job) RemoveJob(ctx *gin.Context) {
	id := ctx.Query("id")

	err := m.jobDAO.RemoveJobFromDb(id)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Job booked successfully"})
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// UpdateJob godoc
// @Summary Update a job
// @Description Update a job
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} models.Message
// @Router /Job/UpdateJob [put]
func (m *Job) UpdateJob(ctx *gin.Context) {

	id := ctx.Query("id")
	var newJob daos.Jobs
	if err := ctx.ShouldBindJSON(&newJob.Job); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(newJob.Job)
	err := m.jobDAO.UpdateJobToDb(newJob, id)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Job booked successfully"})
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}
