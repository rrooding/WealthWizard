package main

import (
	"flag"
	"os"
  "io/ioutil"
	"testing"
  "reflect"
)

func Test_getCliInput(t *testing.T) {
  tests := []struct {
    name string
    want cliInput
    wantErr bool
    osArgs []string
  }{
    {"Default parameters", cliInput{"test.csv", false}, false, []string{"degiro-importer", "test.csv"}},
    {"No parameters", cliInput{}, true, []string{"cmd"}},
    {"Dry run enabled", cliInput{"test.csv", true}, false, []string{"degiro-importer", "-dryRun", "test.csv"}},
    {"Parameter but no filename", cliInput{}, true, []string{"degiro-importer", "-dryRun"}},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      actualOsArgs := os.Args

      defer func() {
        os.Args = actualOsArgs
        flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
      }()

      os.Args = tt.osArgs
      got, err := getCliInput()

      if (err != nil) != tt.wantErr {
        t.Errorf("getCliInput() error = %v, wantErr %v", err, tt.wantErr)
        return
      }

      if !reflect.DeepEqual(got, tt.want) {
        t.Errorf("getCliInput() = %v, want %v", got, tt.want)
      }
    })
  }
}

func Test_checkIfValidFile(t *testing.T) {
  tmpfile, err := ioutil.TempFile("", "test*.csv")
  if err != nil {
    panic(err)
  }

  defer os.Remove(tmpfile.Name())

  tests := []struct {
    name string
    filename string
    want bool
    wantErr bool
  }{
    {"File exists", tmpfile.Name(), true, false},
    {"File does not exist", "doesnotexist.csv", false, true},
    {"File is not CSV", "movie.mp4", false, true},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      got, err := checkIfValidFile(tt.filename)
      if (err != nil) != tt.wantErr {
        t.Errorf("checkIfValidFile() error = %v, wantErr %v", err, tt.wantErr)
        return
      }

      if got != tt.want {
        t.Errorf("checkIfValidFile() = %v, want %v", got, tt.want)
      }
    })
  }
}

func Test_newTransaction(t *testing.T) {
  tests := [] struct{
    name string
    data map[string]string
    want *NewTransaction
    wantErr bool
  }{
    {"Correct data", map[string]string{"ISIN": "123", "Order ID": "456"}, &NewTransaction{Isin: "123", Broker: "DeGiro", BrokerId: "456"}, false},
    {"Correct data without order id", map[string]string{"ISIN": "123"}, &NewTransaction{Isin: "123", Broker: "DeGiro"}, false},
    {"Missing ISIN", map[string]string{"Order ID": "123"}, nil, true},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      get, err := newTransaction(tt.data)
      if (err != nil) != tt.wantErr {
        t.Errorf("newTransaction() error = %v, wantErr %v", err, tt.wantErr)
      }

      if !reflect.DeepEqual(get, tt.want) {
        t.Errorf("newTransaction() = %v, want %v", get, tt.want)
      }
    })
  }
}
