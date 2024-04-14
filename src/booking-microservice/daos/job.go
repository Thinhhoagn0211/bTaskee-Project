package daos

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/databases"
	rabbitmq "github.com/thinhhb0211/job-distribution-system/pricing-microservice/rabbitmq"
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

// Job manages Movie CRUD
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
)

// AddJobToDb adds a new job to the database
func (m *Jobs) AddJobToDb(newJob Jobs) error {

	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(COLLECTION)
	filter := bson.M{"name": newJob.Job.Service}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if count == 0 {
		return fmt.Errorf("no service found with name: %s", newJob.Job.Service)
	}

	_, err = time.Parse(time.RFC3339, newJob.Job.CreatedAt)
	if err != nil {
		{
			return fmt.Errorf("invalid time format: %s", err)
		}
	}

	collection = databases.Database.MgDbClient.Database("companyDatabase").Collection(JOB_ORDERED_COLLECTION)

	filter = bson.M{"id": newJob.Job.ID}
	result, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if result > 0 {
		return fmt.Errorf("job with ID %s already exists", newJob.Job.ID)
	}

	_, err = collection.InsertOne(context.TODO(), newJob.Job)
	if err != nil {
		log.Fatal(err)
		return err
	}

	rabbitmq.RabbitMQConfig.Publish(newJob.Job)
	return nil
}

// RemoveJobFromDb removes a job from the database
func (m *Jobs) RemoveJobFromDb(id string) error {

	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(JOB_ORDERED_COLLECTION)

	filter := bson.M{"id": id}
	result, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if result == 0 {
		return fmt.Errorf("no service found with id: %s", id)
	}

	_, err = collection.DeleteMany(context.TODO(), bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// UpdateJobToDb updates a job in the database
func (m *Jobs) UpdateJobToDb(newJob Jobs, id string) error {

	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(JOB_ORDERED_COLLECTION)

	filter := bson.M{"id": id}
	result, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if result == 0 {
		return fmt.Errorf("no service found with id: %s", id)
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": newJob.Job})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
