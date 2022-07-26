package bank

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(credits float64) error {

	if credits > 0 {
		a.Balance += credits

		return nil
	}

	return errors.New("the amount to deposit should be greater than zero")
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {

		return errors.New("the amount to withdraw should be greater than zero")
	} else if amount > a.Balance {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}
	a.Balance -= amount
	return nil

}

func (a *Account) Statement() string {

	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}
