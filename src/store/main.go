package main

import (
	"fmt"

	"github.com/junior-isabel/store"
)

func main() {

	employee := store.CreateEmployee("José", "Júnior", 2000.00)
	fmt.Println(employee)

	employee.AddCredits(250)
	fmt.Println(employee)

	employee.RemoveCredits(50)
	fmt.Println(employee)

	fmt.Println(employee.CheckCredits())

	employee.ChangeName("Chicapue")
	fmt.Println(employee)
}
