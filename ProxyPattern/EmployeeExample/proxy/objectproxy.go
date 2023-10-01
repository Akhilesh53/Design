package proxy

import "errors"

type EmployeeProxy struct {
	employee IEmployee
}

func NewEmployeeProxy() IEmployee {
	return &EmployeeProxy{
		employee: newEmployee(),
	}
}

func (ep *EmployeeProxy) Create(client string) error {
	if client == "ADMIN" {
		ep.employee.Create(client)
		return nil
	}
	return errors.New(" permission denied")
}

func (ep *EmployeeProxy) Delete(client string) error {
	if client == "ADMIN" {
		ep.employee.Delete(client)
		return nil
	}
	return errors.New(" permission denied")
}

func (ep *EmployeeProxy) EmployeeDo(client string) error {
	if client == "ADMIN" || client == "USER" {
		ep.employee.EmployeeDo(client)
		return nil
	}
	return errors.New(" permission denied")
}
