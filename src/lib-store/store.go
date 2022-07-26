package store

import "fmt"

type Account struct {
	name     string
	nickname string
}

func (account *Account) ChangeName(name string) {
	if len(name) > 0 {
		account.name = name

	}
}

type Employee struct {
	credits float64
	Account
}

func (e *Employee) String() string {

	return fmt.Sprintf("User Account:\n\t%s %s\n\t$%.2f", e.name, e.nickname, e.credits)
}
func CreateEmployee(firstname string, lastname string, credits float64) *Employee {

	return &Employee{
		credits,
		Account{
			name:     firstname,
			nickname: lastname,
		},
	}
}

func (emp *Employee) AddCredits(credits float64) {
	if credits > 0 {
		emp.credits += credits
	}
}

func (emp *Employee) RemoveCredits(credits float64) {
	if credits > 0 && emp.credits >= credits {
		emp.credits -= credits
	}
}

func (emp *Employee) CheckCredits() float64 {

	return emp.credits
}
