package main

import (
	"context"
	"fmt"
	"log"

	"EmployeeLearning/Employee/protos/gen"

	"google.golang.org/grpc"
)

func main() {

	// Implemented only for checking the localhost
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := gen.NewEmployeeServiceClient(conn)

	emp, err := client.CreateEmployee(context.Background(), &gen.CreateEmployeeRequest{
		Employee: &gen.Employee{
			Firstname:   "Sahil",
			Lastname:    "Agrawal",
			Username:    "SahilAgrawal99",
			Gender:      "Male",
			Age:         22,
			Department:  "TTH",
			Designation: "Software Engineer",
			Salary:      25000.00,
		},
	})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Employee Successfully created", emp.Employee)
	}
}
