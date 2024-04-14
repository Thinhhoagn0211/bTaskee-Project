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

// GetAllCurrentService godoc
// @Summary List all current services
// @Description List all current services
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} listService interface{}
// @Router /Job/GetAllServices [get]
func (m *Job) GetAllCurrentService(ctx *gin.Context) {
	listService, err := m.jobDAO.GetAllServices()
	if err == nil {
		ctx.JSON(http.StatusOK, listService)
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{"Cannot retrieve services information"})
		log.Debug("[ERROR]: ", err)
	}
}

// GetAllEmployees godoc
// @Summary List all current employees
// @Description List all current employees
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} listService interface{}
// @Router /Job/GetAllEmployees [get]
func (m *Job) GetAllEmployees(ctx *gin.Context) {
	listService, err := m.jobDAO.GetNameEmployeesFromDb()
	if err == nil {
		ctx.JSON(http.StatusOK, listService)
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{"Cannot retrieve services information"})
		log.Debug("[ERROR]: ", err)
	}
}

// GetAllJobsOrdered godoc
// @Summary List all jobs ordered
// @Description List all jobs ordered
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} listService interface{}
// @Router /Job/GetAllJobsOrdered [get]
func (m *Job) GetAllJobsOrdered(ctx *gin.Context) {
	listService, err := m.jobDAO.GetJobsFromDb()
	if err == nil {
		ctx.JSON(http.StatusOK, listService)
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{"Cannot retrieve services information"})
		log.Debug("[ERROR]: ", err)
	}
}

// GetJobsById godoc
// @Summary List all job by ID
// @Description List all job by ID
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} listService interface{}
// @Router /Job/GetJobsByID [get]
func (m *Job) GetJobsById(ctx *gin.Context) {
	id := ctx.Query("id")
	listService, err := m.jobDAO.GetJobsByIdFromDb(id)
	if err == nil {
		ctx.JSON(http.StatusOK, listService)
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{"Cannot retrieve services information"})
		log.Debug("[ERROR]: ", err)
	}
}

// AddService godoc
// @Summary Add a new service
// @Description Add a new service
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} message string
// @Router /Job/AddService [post]
func (m *Job) AddService(ctx *gin.Context) {

	var newJob daos.Jobs
	if err := ctx.ShouldBindJSON(&newJob.ServiceObj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(newJob.Job)
	err := m.jobDAO.AddServiceToDb(newJob)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Service add successfully"})
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// AddEmployee godoc
// @Summary Add a new service
// @Description Add a new service
// @Tags job
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.Error
// @Success 200 {object} message string
// @Router /Job/AddEmployee [get]
func (m *Job) AddEmployee(ctx *gin.Context) {

	var newJob daos.Jobs
	if err := ctx.ShouldBindJSON(&newJob.Employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(newJob.Job)
	err := m.jobDAO.AddEmployeeToDb(newJob)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Employee add successfully"})
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}
