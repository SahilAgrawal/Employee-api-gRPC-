package controller

import (
	"context"
	"fmt"

	"EmployeeLearning/Employee/server/db"

	"EmployeeLearning/Employee/protos/gen"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmployeeServiceServerImplementation struct {
	gen.UnimplementedEmployeeServiceServer
}

func (c *EmployeeServiceServerImplementation) CreateEmployee(ctx context.Context, in *gen.CreateEmployeeRequest) (*gen.CreateEmployeeResponse, error) {

	employee := in.GetEmployee()
	fmt.Println(employee)

	data := RequestToData(employee)

	res, err := db.InsertEmployee(data)
	if res == nil || err != nil {
		return &gen.CreateEmployeeResponse{}, status.Errorf(codes.Internal, "Failed to create", err)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return &gen.CreateEmployeeResponse{}, status.Errorf(codes.Internal, "INTERNAL SERVER ERROR")
		// log.Fatal(ok)
	}

	data.Id = oid
	return &gen.CreateEmployeeResponse{
		Employee: DataToResponse(data),
	}, nil

}

func (c *EmployeeServiceServerImplementation) UpdateEmployee(ctx context.Context, in *gen.UpdateEmployeeRequest) (*gen.UpdateEmployeeResponse, error) {
	id := in.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &gen.UpdateEmployeeResponse{}, status.Errorf(codes.NotFound, "Error : %v", err)
	}

	employee := in.GetEmployee()
	data := RequestToData(employee)

	updateCount, err := db.UpdateEmployee(objectId, data)

	fmt.Println(updateCount)
	if err != nil {
		return &gen.UpdateEmployeeResponse{}, err
	} else if updateCount == 1 {
		data.Id = objectId
		return &gen.UpdateEmployeeResponse{Employee: DataToResponse(data)}, nil
	}
	fmt.Println("Returning from end")
	return &gen.UpdateEmployeeResponse{}, status.Errorf(codes.Internal, "Internal Server error")
}

func (c *EmployeeServiceServerImplementation) GetEmployeeById(ctx context.Context, in *gen.GetEmployeeByIdRequest) (*gen.GetEmployeeByIdResponse, error) {
	id := in.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &gen.GetEmployeeByIdResponse{}, status.Errorf(codes.NotFound, "Error : %v", err)
	}

	data, err := db.GetEmployeeById(objectId)
	if err != nil {
		return &gen.GetEmployeeByIdResponse{}, status.Errorf(codes.NotFound, "Error : %v ", err)
	}

	return &gen.GetEmployeeByIdResponse{
		Employee: DataToResponse(data),
	}, nil
}

func (c *EmployeeServiceServerImplementation) GetEmployees(ctx context.Context, in *gen.GetEmployeesRequest) (*gen.GetEmployeeResponse, error) {

	employees, err := db.GetEmployees()
	if err != nil {
		return &gen.GetEmployeeResponse{}, err
	}

	var result []*gen.Employee
	for _, employee := range employees {
		result = append(result, DataToResponse(employee))
	}

	return &gen.GetEmployeeResponse{
		Employee: result,
	}, nil
}

func (c *EmployeeServiceServerImplementation) DeleteEmployee(ctx context.Context, in *gen.DeleteEmployeeByIdRequest) (*gen.DeleteEmployeeByIdResponse, error) {
	id := in.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &gen.DeleteEmployeeByIdResponse{}, status.Errorf(codes.NotFound, "Error : %v", err)
	}

	deleteCount, err := db.DeleteEmployee(objectId)

	if err != nil {
		return &gen.DeleteEmployeeByIdResponse{}, err
	} else if deleteCount == 1 {
		return &gen.DeleteEmployeeByIdResponse{
			DeleteCount: deleteCount,
		}, nil
	}
	return &gen.DeleteEmployeeByIdResponse{
		DeleteCount: 0,
	}, status.Errorf(codes.Internal, "Internal Server error")
}
