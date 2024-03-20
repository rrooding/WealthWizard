package transaction

import (
	"fmt"

	"wealth-wizard/degiro-importer/api"
)

type TransactionCreator interface {
	CreateTransaction(api.NewTransaction) (*api.CreateTransactionResponse, error)
}

func CreatorFunc(client TransactionCreator) func(*api.NewTransaction) {
	return func(t *api.NewTransaction) {
		Println(t)

		res, err := client.CreateTransaction(*t)
		if err != nil {
			fmt.Printf("Error creating transaction: %v", err)
			return
		}

		fmt.Printf("Transaction created: %v\n", res.GetCreateTransaction().Id)
	}
}
