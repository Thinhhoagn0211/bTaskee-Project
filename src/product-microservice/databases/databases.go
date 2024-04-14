package databases

import (
	"context"

	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDbClient   *mongo.Client
	Databasename string
}

// Init initializes mongo database
func (db *MongoDB) Init() error {
	uri := "mongodb://" + common.Config.MgAddrs
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	client.Database("metadata")
	db.MgDbClient = client
	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	db.MgDbClient.Disconnect(context.Background())
}
