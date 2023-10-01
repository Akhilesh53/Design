package proxy

import "fmt"

type IEmployee interface {
	Create(string) error
	Delete(string) error
	EmployeeDo(string) error
}

type Employee struct {
}

func newEmployee() IEmployee {
return &Employee{}
}

func (e Employee) Create(client string) error {
	fmt.Println("Creating Employee by client: ", client)
	return nil
}

func (e Employee) Delete(client string) error {
	fmt.Println("Deleting Employee by client: ", client)
	return nil
}

func (e Employee) EmployeeDo(client string) error {
	fmt.Println("Employee Do by client: ", client)
	return nil
}
