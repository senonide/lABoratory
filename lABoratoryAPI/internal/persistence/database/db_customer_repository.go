package database

import (
	"context"
	"fmt"
	"lABoratory/lABoratoryAPI/internal/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CustomersCollName string = "customers"

type dbCustomerRepository struct {
	database *mongo.Database
}

func NewDbCustomerRepository() *dbCustomerRepository {
	repository := new(dbCustomerRepository)
	repository.database = GetDatabase()
	return repository
}

func (r *dbCustomerRepository) GetAll(experimentId string) ([]models.Customer, error) {
	ctx := context.Background()
	collection := r.database.Collection(CustomersCollName)
	customers := []models.Customer{}
	filter := bson.M{"experimentid": experimentId}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var customer models.Customer
		err = cur.Decode(&customer)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (r *dbCustomerRepository) GetOne(customerKey, experimentId string) (*models.Customer, error) {
	ctx := context.Background()
	collection := r.database.Collection(CustomersCollName)
	var customer *models.Customer
	filter := bson.M{"key": customerKey, "experimentid": experimentId}
	cur := collection.FindOne(ctx, filter)
	if cur.Err() != nil {
		return nil, nil
	}
	err := cur.Decode(&customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *dbCustomerRepository) Create(customer models.Customer) (string, error) {
	ctx := context.Background()
	collection := r.database.Collection(CustomersCollName)
	result, err := collection.InsertOne(ctx, customer)
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("%v", result)
	id = strings.Replace(id, "&{ObjectID(\"", "", 1)
	id = strings.Replace(id, "\")}", "", 1)
	return id, nil
}

func (r *dbCustomerRepository) SetAssignment(customerKey, experimentId string, newAssigment models.Assignment) error {
	ctx := context.Background()
	collection := r.database.Collection(CustomersCollName)
	filter := bson.M{"key": customerKey, "experimentid": experimentId}
	update := bson.M{
		"$set": bson.M{
			"assignmentname":        newAssigment.AssignmentName,
			"assignmentdescription": newAssigment.AssignmentDescription,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *dbCustomerRepository) SetAllAssignments(experimentId string, newAssigment models.Assignment) error {
	ctx := context.Background()
	collection := r.database.Collection(CustomersCollName)
	filter := bson.M{"experimentid": experimentId}
	update := bson.M{
		"$set": bson.M{
			"assignmentname":        newAssigment.AssignmentName,
			"assignmentdescription": newAssigment.AssignmentDescription,
		},
	}
	_, err := collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *dbCustomerRepository) DeleteAll(experimentId string) (bool, error) {
	ctx := context.Background()
	collection := r.database.Collection(CustomersCollName)
	filter := bson.M{"experimentid": experimentId}
	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}
