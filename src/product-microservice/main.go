package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/common"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/controllers"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/databases"
	_ "github.com/thinhhb0211/job-distribution-system/pricing-microservice/docs"
)

// Main manages main golang application
type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize mongo database
	err = databases.Database.Init()
	if err != nil {
		return err
	}

	m.router = gin.Default()

	return nil
}

func main() {
	m := Main{}

	// Initialize server
	if m.initServer() != nil {
		return
	}

	defer databases.Database.Close()

	c := controllers.Job{}

	// Simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		// Job Service
		v1.GET("/Job/GetAllServices", c.GetAllCurrentService)
		v1.POST("/Job/AddService", c.AddService)
		v1.POST("/Job/AddEmployee", c.AddEmployee)
		v1.GET("/Job/GetAllEmployees", c.GetAllEmployees)
		v1.GET("/Job/GetAllJobsOrdered", c.GetAllJobsOrdered)
		v1.GET("/Job/GetJobsByID", c.GetJobsById)

	}
	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	m.router.Run(common.Config.Port)
}
