package daos

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/databases"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type JobAttr struct {
	ID           string `json:"ID" bson:"ID"`
	NameEmployee string `json:"nameEmployee" bson:"nameEmployee"`
	CreatedAt    string `json:"createdAt" bson:"createdAt"`
	Service      string `json:"service" bson:"service"`
	Note         string `json:"note" bson:"note"`
}

type EmployeeAttr struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Movie manages Movie CRUD
type Jobs struct {
	ListServices []bson.M `bson:"name"`
	Employee     EmployeeAttr
	NameService  []bson.M `bson:"service"`
	ServiceObj   Service
	Job          JobAttr
}

// Create a slice to store the results
type totalPrice struct {
	TotalPrice float64 `bson:"totalPrice"`
}

// COLLECTION of the database table
const (
	COLLECTION             = "productService"
	EMPLOYEE_COLLECTION    = "ListOfEmployee"
	JOB_ORDERED_COLLECTION = "JobsOrdered"
	LIST_JOB               = "list_job"
)

// CalculatePriceJobs calculates the total price of all jobs
func (m *Jobs) CalculatePriceJobs() (int, error) {

	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(COLLECTION)

	findOptions := options.Find().SetProjection(bson.M{"_id": 0})

	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	var priceProduct []map[string]int
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var result Jobs
		err := cursor.Decode(&result.ServiceObj)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		priceProduct = append(priceProduct, map[string]int{result.ServiceObj.Name: result.ServiceObj.Price})
	}

	collection = databases.Database.MgDbClient.Database("companyDatabase").Collection(JOB_ORDERED_COLLECTION)

	findOptions = options.Find().SetProjection(bson.M{"_id": 0})

	cursor, err = collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer cursor.Close(context.TODO())

	var totalPrice int
	for cursor.Next(context.TODO()) {
		var result Jobs
		err := cursor.Decode(&result.Job)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		for _, product := range priceProduct {
			for key, value := range product {
				if key == result.Job.Service {
					totalPrice += value
					break
				}

			}
		}
	}

	return totalPrice, nil
}

// CalculatePriceJobsByTimeFromDb calculates the total price of all jobs in the database
func (m *Jobs) CalculatePriceJobsByTimeFromDb(startTime string, endTime string) (int, error) {

	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(COLLECTION)

	findOptions := options.Find().SetProjection(bson.M{"_id": 0})

	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	var priceProduct []map[string]int
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var result Jobs
		err := cursor.Decode(&result.ServiceObj)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		priceProduct = append(priceProduct, map[string]int{result.ServiceObj.Name: result.ServiceObj.Price})
	}

	collection = databases.Database.MgDbClient.Database("companyDatabase").Collection(JOB_ORDERED_COLLECTION)

	findOptions = options.Find().SetProjection(bson.M{"_id": 0})

	cursor, err = collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer cursor.Close(context.TODO())
	var totalPrice int
	for cursor.Next(context.TODO()) {
		var result Jobs
		err := cursor.Decode(&result.Job)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		fmt.Println("Job CreatedAt:", result.Job.CreatedAt)

		jobTime, err := time.Parse(time.RFC3339, result.Job.CreatedAt)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}

		start, err := time.Parse(time.RFC3339, startTime)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}

		end, err := time.Parse(time.RFC3339, endTime)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}

		for _, product := range priceProduct {
			for key, value := range product {
				if key == result.Job.Service && jobTime.After(start) && jobTime.Before(end) {
					totalPrice += value
					break
				}
			}
		}
	}

	return totalPrice, nil
}
