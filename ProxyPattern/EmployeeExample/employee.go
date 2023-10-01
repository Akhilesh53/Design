package employeeexample

import (
	"fmt"
	"pattern/ProxyPattern/EmployeeExample/proxy"
)

func CallProxyPattern() {
	var employee = proxy.NewEmployeeProxy()

	err := employee.Create("test")
	if err != nil {
		fmt.Println(err)
	}

	err = employee.Create("ADMIN")
	if err != nil {
		fmt.Println(err)
	}

	err = employee.Delete("test")
	if err != nil {
		fmt.Println(err)
	}

	err = employee.Delete("ADMIN")
	if err != nil {
		fmt.Println(err)
	}

	err = employee.EmployeeDo("test")
	if err != nil {
		fmt.Println(err)
	}

	err = employee.EmployeeDo("ADMIN")
	if err != nil {
		fmt.Println(err)
	}
}
