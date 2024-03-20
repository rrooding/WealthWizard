package transaction

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"wealth-wizard/degiro-importer/api"
)

func newTransaction(data map[string]string) (*api.NewTransaction, error) {
	if data["ISIN"] == "" {
		return nil, errors.New("ISIN is required")
	}

	// Parse the amount as an integer
	amount, err := strconv.Atoi(data["Aantal"])
	if err != nil {
		return nil, errors.New("Invalid amount")
	}

	// Parse the date and time
	dateTime := fmt.Sprintf("%s %s CET", data["Datum"], data["Tijd"])
	date, err := time.Parse("02-01-2006 15:04 MST", dateTime)
	if err != nil {
		return nil, err
	}

	// Parse the price as a decimal
	price := api.MoneyInput{
		Amount:   data["Koers"],
		Currency: data["8"],
	}

	newTransaction := &api.NewTransaction{
		Isin:     data["ISIN"],
		Broker:   "DeGiro",
		Date:     date,
		Exchange: data["Beurs"],
		Amount:   amount,
		Price:    price,
	}

	// Fill in the broker ID if it exists
	if data["Order ID"] != "" {
		newTransaction.BrokerId = data["Order ID"]
	}

	if data["Transactiekosten en/of"] != "" {
		// Parse the transaction cost as a decimal
		newTransaction.TransactionCost = api.MoneyInput{
			Amount:   data["Transactiekosten en/of"],
			Currency: data["15"],
		}
	}

	return newTransaction, nil
}

func HandleTransaction(writerChannel <-chan map[string]string, done chan<- bool, callback func(*api.NewTransaction)) {
	for {
		record, ok := <-writerChannel

		if ok {
			t, err := newTransaction(record)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}

			callback(t)
		} else {
			done <- true
			break
		}
	}
}

func Println(t *api.NewTransaction) {
	fmt.Printf("Transaction: %v\n", t)
}
