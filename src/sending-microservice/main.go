package main

import (
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/common"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/databases"
	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/rabbitmq"
)

// Main manages main golang application
type Main struct {
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

	return nil
}

func main() {
	m := Main{}

	// Initialize server
	if m.initServer() != nil {
		return
	}

	defer databases.Database.Close()
	var r rabbitmq.RabbitMQ
	r.Consume()
}
