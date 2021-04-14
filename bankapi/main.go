package main

import (
	"encoding/json"
	"fmt"
	"github.com/leeranzhi/GoLearn/bankcore"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[float64]*CustomAccount{}

func main() {
	accounts[1001] = &CustomAccount{
		Account: &bankcore.Account{
			Customer: bankcore.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "086 5550166",
			},
			Number:  1001,
			Balance: 0,
		},
	}

	accounts[1002] = &CustomAccount{
		Account: &bankcore.Account{
			Customer: bankcore.Customer{
				Name:    "Mark",
				Address: "Los Angeles, California",
				Phone:   "086 5550166",
			},
			Number:  1002,
			Balance: 0,
		},
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
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
			json.NewEncoder(w).Encode(bankcore.Statement(account))
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

func transfer(writer http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")
	destqs := req.URL.Query().Get("dest")

	if numberqs == "" {
		fmt.Fprintf(writer, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(writer, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(writer, "Invalid amount number!%v", err)
	} else if dest, err := strconv.ParseFloat(destqs, 64); err != nil {
		fmt.Fprintf(writer, "Invalid amount number!%v", err)
	} else {
		if accountA, ok := accounts[number]; !ok {
			fmt.Fprintf(writer, "Account with number %v can't be found!", number)
		} else if accountB, ok := accounts[dest]; !ok {
			fmt.Fprintf(writer, "Account with number %v can't be found!", destqs)
		} else {
			err := accountA.Transfer(amount, accountB.Account)
			if err != nil {
				fmt.Fprintf(writer, "%v", err)
			} else {
				fmt.Fprintf(writer, accountA.Statement())
			}
		}
	}
}

type CustomAccount struct {
	*bankcore.Account
}

func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(json)
}
