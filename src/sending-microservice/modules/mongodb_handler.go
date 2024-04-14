package modules

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/databases"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type JobAttr struct {
	ID        string `json:"id" bson:"id"`
	CreatedAt string `json:"createdAt" bson:"createdAt"`
	Service   string `json:"service" bson:"service"`
	Note      string `json:"note" bson:"note"`
}

type EmployeeAttr struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Movie manages Movie CRUD
type Jobs struct {
	Employee   EmployeeAttr
	ServiceObj Service
	Job        JobAttr
}

// COLLECTION of the database table
const (
	COLLECTION             = "productService"
	EMPLOYEE_COLLECTION    = "ListOfEmployee"
	JOB_ORDERED_COLLECTION = "JobsOrdered"
	LIST_JOB               = "list_job"
	DATABASE               = "companyDatabase"
)

// Insert adds a new Movie into database'
func (m *Jobs) InsertEmployeeToService(newJob Jobs) error {

	collection := databases.Database.MgDbClient.Database(DATABASE).Collection(EMPLOYEE_COLLECTION)

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	var listEmployee []string
	for cursor.Next(context.Background()) {
		var result Jobs
		err := cursor.Decode(&result.Employee)
		if err != nil {
			log.Fatal(err)
			return err
		}
		listEmployee = append(listEmployee, result.Employee.Name)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(listEmployee))

	// Get the value at the random index
	randomValue := listEmployee[randomIndex]
	fmt.Println(randomValue)

	collection = databases.Database.MgDbClient.Database(DATABASE).Collection(JOB_ORDERED_COLLECTION)
	_, err = collection.UpdateOne(context.TODO(), bson.M{"id": newJob.Job.ID}, bson.M{"$set": bson.M{"employee": randomValue}})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
