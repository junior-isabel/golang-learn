package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/junior-isabel/bank"
)

var accounts = map[float64]*bank.Account{}

func main() {

	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	log.Fatal(http.ListenAndServe("localhost:3333", nil))

}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		if account, ok := accounts[number]; !ok {
			fmt.Fprintf(w, "User not exists!")
		} else {
			fmt.Fprintln(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	number, err := strconv.ParseFloat(numberqs, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid account number!")
		return
	} else if amount, errAmount := strconv.ParseFloat(amountqs, 64); errAmount != nil {
		fmt.Fprintf(w, "Amount number invalid!\n")
		return
	} else {

		if acount, ok := accounts[number]; !ok {

			fmt.Fprintf(w, "user not exists")
			return

		} else {
			err := acount.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "do Not allow deposit %v\n", amount)
			} else {
				fmt.Fprintf(w, "Amount deposit success!\n")

			}
		}
	}

}
