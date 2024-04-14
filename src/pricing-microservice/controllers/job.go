package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/common"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/daos"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/models"
)

// Job manages Job CRUD
type Job struct {
	jobDAO daos.Jobs
}

// CalculatePriceJobs godoc
// @Summary Calculate the total price of all jobs
// @Description Calculate the total price of all jobs
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} totalPrice
// @Router /CalculatePrice/AllJobs [get]
func (m *Job) CalculatePriceJobs(ctx *gin.Context) {

	totalPrice, err := m.jobDAO.CalculatePriceJobs()

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"totalPriceOfAllJobs": totalPrice})
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{common.StatusCodeUnknown, "Cannot retrieve movie information"})
		log.Debug("[ERROR]: ", err)
	}
}

// CalculatePriceJobsByTime godoc
// @Summary Calculate the total price of all jobs by time
// @Description Calculate the total price of all jobs by time
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} totalPriceByTime
// @Router /CalculatePrice/JobByTime [get]
func (m *Job) CalculatePriceJobsByTime(ctx *gin.Context) {

	startTime := ctx.Query("startTime")
	endTime := ctx.Query("endTime")

	totalPrice, err := m.jobDAO.CalculatePriceJobsByTimeFromDb(startTime, endTime)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"totalPriceByTime": totalPrice})
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{common.StatusCodeUnknown, "Cannot retrieve movie information"})
		log.Debug("[ERROR]: ", err)
	}
}
