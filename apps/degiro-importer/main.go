package main

import (
  "errors"
  "flag"
  "fmt"
  "os"
  "io"
  "encoding/csv"
  "path/filepath"
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

func processCsvFile(input cliInput, writerChannel chan<- map[string]string) {
  // Open the file
  file, err := os.Open(input.filepath)
  check(err)
  defer file.Close()

  // Define headers and line slice
  var headers, line []string

  // Initialize the CSV reader
  reader := csv.NewReader(file)

  // Reading the first line, where we will find our headers
  headers, err = reader.Read()
  check(err)

  // Iterate over each line
  for {
    line, err = reader.Read()
    if err == io.EOF {
      close(writerChannel)
      break
    } else if err != nil {
      exitGracefully(err)
    }

    record, err := processCsvLine(headers, line)

    if err != nil {
      fmt.Printf("Line: %sError: %s\n", line, err)
      continue
    }

    writerChannel <- record
  }
}

func processCsvLine(headers []string, dataList []string) (map[string]string, error) {
  if len(dataList) != len(headers) {
    return nil, errors.New("Line doesn't match headers format")
  }

  recordMap := make(map[string]string)
  for i, name := range headers {
    recordMap[name] = dataList[i]
  }

  return recordMap, nil
}

/******/

type transaction struct {
  ISIN string
  Broker string
  LocalId string
}

func transactionForMap(data map[string]string) (*transaction, error) {
  return &transaction{
    ISIN: data["ISIN"],
    Broker: "DeGiro",
    LocalId: data["Order ID"],
  }, nil
}

/******/

func handleTransaction(writerChannel <-chan map[string]string, done chan<- bool) {
  for {
    record, ok := <- writerChannel

    if ok {
      t, err := transactionForMap(record)
      check(err)
      fmt.Printf("Line %v\n", t)

    } else {
      done <- true
      break
    }
  }
}

func main() {
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
