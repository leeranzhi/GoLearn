package main

import (
	"fmt"
	"github.com/leeranzhi/GoLearn/bankcore"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[float64]*bankcore.Account{}

func main() {
	accounts[1001] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "086 5550166",
		},
		Number:  1001,
		Balance: 0,
	}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
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
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with %v can't be found!", number)
		} else {
			fmt.Fprintf(w, account.Statement())
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
	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!%v", err)
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!%v", err)
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}
