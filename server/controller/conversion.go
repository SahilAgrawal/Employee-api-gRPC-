package controller

import (
	"EmployeeLearning/Employee/protos/gen"
	"EmployeeLearning/Employee/server/model"
)

func RequestToData(employee *gen.Employee) model.Employee {
	return model.Employee{
		Firstname:   employee.GetFirstname(),
		Lastname:    employee.GetLastname(),
		Username:    employee.GetUsername(),
		Gender:      employee.GetGender(),
		Age:         employee.GetAge(),
		Department:  employee.GetDepartment(),
		Designation: employee.GetDesignation(),
		Salary:      employee.GetSalary(),
	}
}

func DataToResponse(data model.Employee) *gen.Employee {
	return &gen.Employee{
		Id:          data.Id.Hex(),
		Firstname:   data.Firstname,
		Lastname:    data.Lastname,
		Username:    data.Username,
		Gender:      data.Gender,
		Age:         data.Age,
		Department:  data.Department,
		Designation: data.Designation,
		Salary:      data.Salary,
	}
}
