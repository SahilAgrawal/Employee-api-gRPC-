package db

import (
	"context"
	"fmt"
	"log"

	"EmployeeLearning/Employee/server/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

// Insert MongoURL, Database and Collection name
const uri = "<MONGO_URI>"
const dbName = "<DB_NAME>"
const colName = "<COLLECTION_NAME>"

var collection *mongo.Collection
var client *mongo.Client

func ConnectDB() {

	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithCancel(context.Background())

	err = client.Connect(ctx)
	// defer client.Disconnect(ctx)

	if err != nil {
		fmt.Println("Getting trouble to connect the DB")
		log.Fatal(err)

	}

	// client, err := mongo.Connect(context.TODO(), clientOption)

	fmt.Println("Mongo DB Connection Success...")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Connection Instance ready..")

}

func InsertEmployee(employee model.Employee) (*mongo.InsertOneResult, error) {
	inserted, err := collection.InsertOne(context.TODO(), employee)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	fmt.Println("Insert Data Id ", inserted.InsertedID)
	return inserted, err

}

func GetEmployeeById(objectId primitive.ObjectID) (model.Employee, error) {

	var result model.Employee
	filter := bson.M{"_id": objectId}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		// log.Fatal(err)
		return result, err
	}
	// fmt.Println("Insert Data Id ", inserted.InsertedID)
	fmt.Println("The data is ", result)

	return result, nil
}

func GetEmployees() ([]model.Employee, error) {
	var employees []model.Employee

	filter := bson.M{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		// log.Fatal(err)
		return employees, status.Errorf(codes.Internal, "INTERNAL SERVER ERROR")
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var employee model.Employee
		err := cur.Decode(&employee)
		if err != nil {
			return []model.Employee{}, status.Errorf(codes.Internal, err.Error())
		}
		employees = append(employees, employee)
	}

	if err := cur.Err(); err != nil {
		return []model.Employee{}, status.Errorf(codes.Internal, err.Error())
	}

	return employees, nil
}

func UpdateEmployee(id primitive.ObjectID, employee model.Employee) (int64, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"firstname":   employee.Firstname,
		"lastname":    employee.Lastname,
		"username":    employee.Username,
		"gender":      employee.Gender,
		"age":         employee.Age,
		"department":  employee.Department,
		"designation": employee.Designation,
		"salary":      employee.Salary,
	}}

	res, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return 0, status.Errorf(codes.Internal, "Failed to update %v", err)
	}

	return res.MatchedCount, nil

}

func DeleteEmployee(objectID primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": objectID}
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "Failed to delete %v", err)
	}

	return res.DeletedCount, nil
}
