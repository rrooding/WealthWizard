package transaction

import (
  "errors"
  "fmt"
  "os"

	"wealth-wizard/degiro-importer/api"
)

func newTransaction(data map[string]string) (*api.NewTransaction, error) {
  if data["ISIN"] == "" {
    return nil, errors.New("ISIN is required")
  }

  newTransaction := &api.NewTransaction{
    Isin: data["ISIN"],
    Broker: "DeGiro",
  }

  if data["Order ID"] != "" {
    newTransaction.BrokerId = data["Order ID"]
  }

  return newTransaction, nil
}

func HandleTransaction(writerChannel <-chan map[string]string, done chan<- bool, callback func(*api.NewTransaction)) {
  for {
    record, ok := <- writerChannel

    if ok {
      t, err := newTransaction(record)
      if (err != nil) {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
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

func CreatorFunc(client *api.API) func(*api.NewTransaction) {
  return func(t *api.NewTransaction) {
    Println(t)

    res, err := client.CreateTransaction(*t)
    if (err != nil) {
      fmt.Printf("Error creating transaction: %v", err)
      return
    }

    fmt.Printf("Transaction created: %v\n", res.GetCreateTransaction().BrokerId)
  }
}
