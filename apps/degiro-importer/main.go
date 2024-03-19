package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"wealth-wizard/degiro-importer/api"
	"wealth-wizard/degiro-importer/csv"
	"wealth-wizard/degiro-importer/transaction"
)

type cliInput struct {
	filepath string
	dryRun   bool
}

func getCliInput() (cliInput, error) {
	dryRun := flag.Bool("dryRun", false, "Dry run")
	flag.Parse()

	if len(flag.Args()) < 1 {
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

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	input, err := getCliInput()
	if err != nil {
		exitGracefully(err)
	}

	if _, err := checkIfValidFile(input.filepath); err != nil {
		exitGracefully(err)
	}

	writerChannel := make(chan map[string]string)
	done := make(chan bool)

	// In dry run mode, we just print the transactions
	callback := transaction.Println

	// If not in dry run mode, we create the transactions
	if !input.dryRun {
		client := api.Init(os.Getenv("WEALTHWIZARD_API"))
		ok, err := client.IsOK()
		if err != nil || !ok {
			exitGracefully(err)
		}

		callback = transaction.CreatorFunc(client)
	}

	go csv.ProcessFile(input.filepath, writerChannel)
	go transaction.HandleTransaction(writerChannel, done, callback)

	// Wait for a signal that we are done
	<-done
}
