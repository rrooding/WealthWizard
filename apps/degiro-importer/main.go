package main

import (
  "errors"
  "flag"
  "fmt"
  "os"
  "path/filepath"
  "context"
  "net/http"
  "github.com/Khan/genqlient/graphql"
  "github.com/joho/godotenv"
)

type cliInput struct {
  filepath string
  dryRun bool
}

func getCliInput() (cliInput, error) {
  dryRun := flag.Bool("dryRun", false, "Dry run")
  flag.Parse()

  if (len(flag.Args()) < 1) {
    return cliInput{}, errors.New("Please provide a filepath")
  }

  filepath := flag.Arg(0)

  return cliInput{filepath, *dryRun}, nil
}

func checkIfValidFile(filename string) (bool, error) {
  if fileExtension := filepath.Ext(filename); fileExtension != ".csv" {
    return false, fmt.Errorf("File %s is not CSV", filename)
  }

  if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
    return false, fmt.Errorf("File %s does not exist", filename)
  }

  return true, nil
}

func exitGracefully(e error) {
  fmt.Fprintf(os.Stderr, "error: %v\n", e)
  os.Exit(1)
}

func check(e error) {
  if e != nil {
    exitGracefully(e)
  }
}

func newTransaction(data map[string]string) (*NewTransaction, error) {
  if data["ISIN"] == "" {
    return nil, errors.New("ISIN is required")
  }

  newTransaction := &NewTransaction{
    Isin: data["ISIN"],
    Broker: "DeGiro",
  }

  if data["Order ID"] != "" {
    newTransaction.BrokerId = data["Order ID"]
  }

  return newTransaction, nil
}

func handleTransaction(writerChannel <-chan map[string]string, done chan<- bool) {
  ctx := context.Background()
  client := graphql.NewClient("http://localhost:3001/query", http.DefaultClient)

  for {
    record, ok := <- writerChannel

    if ok {
      t, err := newTransaction(record)
      check(err)
      fmt.Printf("newLine %v\n", t)
      resp, err := CreateTransaction(ctx, client, *t)
      check(err)
      fmt.Printf("resp %v\n", resp)

    } else {
      done <- true
      break
    }
  }
}

func main() {
  err := godotenv.Load()
  if (err != nil) {
    fmt.Println("Error loading .env file")
    os.Exit(1)
  }

  flag.Usage = func() {
    fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
    flag.PrintDefaults()
  }

  input, err := getCliInput()
  if (err != nil) {
    exitGracefully(err)
  }

  if _, err := checkIfValidFile(input.filepath); err != nil {
    exitGracefully(err)
  }

  writerChannel := make(chan map[string]string)
  done := make(chan bool)

  go processCsvFile(input, writerChannel)
  go handleTransaction(writerChannel, done)

  // Wait for a signal that we are done
  <- done
}
