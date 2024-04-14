package daos

import (
	"context"
	"fmt"
	"log"

	"github.com/thinhhb0211/job-distribution-system/pricing-microservice/databases"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Service information
type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// Job information
type JobAttr struct {
	ID           string `json:"id" bson:"id"`
	NameEmployee string `json:"nameEmployee" bson:"nameEmployee"`
	CreatedAt    string `json:"createdAt" bson:"createdAt"`
	Service      string `json:"service" bson:"service"`
	Note         string `json:"note" bson:"note"`
}

// Employee information
type EmployeeAttr struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Movie manages Movie CRUD
type Jobs struct {
	Employee    EmployeeAttr
	NameService []bson.M `bson:"service"`
	ServiceObj  Service
	Job         JobAttr
}

// COLLECTION of the database table
const (
	COLLECTION             = "productService"
	EMPLOYEE_COLLECTION    = "ListOfEmployee"
	JOB_ORDERED_COLLECTION = "JobsOrdered"
	LIST_JOB               = "list_job"
)

// GetAllServices gets the list of Service
func (m *Jobs) GetAllServices() (interface{}, error) {

	// Get a collection to execute thbcegmorste query against.
	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(COLLECTION)

	findOptions := options.Find().SetProjection(bson.M{"_id": 0})

	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	listService := make(map[string]interface{})

	var arrayService []any
	for cursor.Next(context.TODO()) {
		var result Jobs
		err = cursor.Decode(&result.ServiceObj)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		arrayService = append(arrayService, result.ServiceObj)
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	listService["Current Services"] = arrayService
	return listService, nil
}

// GetNameEmployeesFromDb gets the list of employees
func (m *Jobs) GetNameEmployeesFromDb() (interface{}, error) {

	// Get a collection to execute the query against.
	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(EMPLOYEE_COLLECTION)

	findOptions := options.Find().SetProjection(bson.M{"_id": 0})

	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	listService := make(map[string]interface{})

	var arrayService []any
	for cursor.Next(context.TODO()) {
		var result Jobs
		err = cursor.Decode(&result.Employee)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		arrayService = append(arrayService, result.Employee)
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	listService["Current Employees"] = arrayService
	return listService, nil
}

// GetJobsFromDb gets the list of jobs from database
func (m *Jobs) GetJobsFromDb() (interface{}, error) {

	// Get a collection to execute the query against.
	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(JOB_ORDERED_COLLECTION)

	findOptions := options.Find().SetProjection(bson.M{"_id": 0})

	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var results Jobs
	if err = cursor.All(context.TODO(), &results.NameService); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return results.NameService, nil
}

// GetJobsByIdFromDb gets the job by id from database
func (m *Jobs) GetJobsByIdFromDb(id string) (interface{}, error) {

	// Get a collection to execute the query against.
	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(JOB_ORDERED_COLLECTION)

	findOptions := options.Find().SetProjection(bson.M{"_id": 0})

	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var results Jobs
	if err = cursor.All(context.TODO(), &results.NameService); err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, job := range results.NameService {
		if job["id"] == id {
			return job, nil
		}
	}
	return nil, err
}

// Insert adds a new Movie into database'
func (m *Jobs) AddServiceToDb(newJob Jobs) error {

	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(COLLECTION)
	filter := bson.M{"name": newJob.ServiceObj.Name}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if count > 0 {
		return fmt.Errorf("service with name %s already exists", newJob.ServiceObj.Name)
	}
	filter = bson.M{"id": newJob.ServiceObj.ID}
	result, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if result > 0 {
		return fmt.Errorf("job with ID %s already exists", newJob.ServiceObj.ID)
	}

	collection = databases.Database.MgDbClient.Database("companyDatabase").Collection(COLLECTION)

	_, err = collection.InsertOne(context.TODO(), newJob.ServiceObj)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Insert adds a new Movie into database'
func (m *Jobs) AddEmployeeToDb(newJob Jobs) error {

	collection := databases.Database.MgDbClient.Database("companyDatabase").Collection(EMPLOYEE_COLLECTION)
	filter := bson.M{"name": newJob.Employee.Name}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if count > 0 {
		return fmt.Errorf("employee with name %s already exists", newJob.Employee.Name)
	}

	collection = databases.Database.MgDbClient.Database("companyDatabase").Collection(EMPLOYEE_COLLECTION)

	filter = bson.M{"id": newJob.Employee.ID}
	result, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if result > 0 {
		return fmt.Errorf("job with ID %s already exists", newJob.Employee.ID)
	}

	_, err = collection.InsertOne(context.TODO(), newJob.Employee)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
