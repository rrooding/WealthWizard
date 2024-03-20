package transaction

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"wealth-wizard/degiro-importer/api"
)

func Test_newTransaction(t *testing.T) {
	cet, _ := time.LoadLocation("CET")
	date := time.Date(2003, 02, 01, 16, 01, 0, 0, cet)

	price := api.MoneyInput{Amount: "34.54", Currency: "EUR"}
	transactionCost := api.MoneyInput{Amount: "12.34", Currency: "EUR"}

	tests := []struct {
		name    string
		data    map[string]string
		want    *api.NewTransaction
		wantErr bool
	}{
		{"Valid data", map[string]string{"ISIN": "123", "Order ID": "456", "Aantal": "1", "Datum": "01-02-2003", "Tijd": "16:01", "Koers": "34.54", "8": "EUR"}, &api.NewTransaction{Isin: "123", Broker: "DeGiro", BrokerId: "456", Amount: 1, Date: date, Price: price}, false},
		{"Valid data with transaction cost", map[string]string{"ISIN": "123", "Order ID": "456", "Aantal": "1", "Datum": "01-02-2003", "Tijd": "16:01", "Koers": "34.54", "8": "EUR", "Transactiekosten en/of": "12.34", "15": "EUR"}, &api.NewTransaction{Isin: "123", Broker: "DeGiro", BrokerId: "456", Amount: 1, Date: date, Price: price, TransactionCost: transactionCost}, false},
		{"Valid data without order id", map[string]string{"ISIN": "123", "Aantal": "1", "Datum": "01-02-2003", "Tijd": "16:01"}, &api.NewTransaction{Isin: "123", Broker: "DeGiro", Amount: 1, Date: date}, false},
		{"Missing ISIN", map[string]string{"Order ID": "456", "Aantal": "1", "Datum": "01-02-2003", "Tijd": "16:01"}, nil, true},
		{"Invalid amount", map[string]string{"ISIN": "123", "Order ID": "456", "Aantal": "test", "Datum": "01-02-2003", "Tijd": "16:01"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTransaction(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("newTransaction() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_HandleTransaction(t *testing.T) {
	mockCallback := func(transaction *api.NewTransaction) {
		fmt.Println("Mock callback called with transaction:", transaction)
	}

	tests := []struct {
		name            string
		writerChannel   <-chan map[string]string
		done            chan bool
		callback        func(*api.NewTransaction)
		expectedOutputs []string
		wantErr         bool
	}{
		{
			name: "Valid data",
			writerChannel: func() <-chan map[string]string {
				ch := make(chan map[string]string)
				go func() {
					ch <- map[string]string{"ISIN": "123", "Order ID": "456", "Aantal": "1", "Koers": "1.23", "8": "EUR", "Datum": "01-02-2003", "Tijd": "16:01"}
					close(ch)
				}()
				return ch
			}(),
			done:            make(chan bool),
			callback:        mockCallback,
			expectedOutputs: []string{"Mock callback called with transaction: &{123 DeGiro 2003-02-01 16:01:00 +0100 CET  1 {1.23 EUR} { } 456}"},
			wantErr:         false,
		},
		{
			name: "Invalid data",
			writerChannel: func() <-chan map[string]string {
				ch := make(chan map[string]string)
				go func() {
					ch <- map[string]string{"Order ID": "456"}
					close(ch)
				}()
				return ch
			}(),
			done:            make(chan bool),
			callback:        mockCallback,
			expectedOutputs: []string{"error: ISIN is required"},
			wantErr:         true,
		},
		{
			name: "Multiple transactions",
			writerChannel: func() <-chan map[string]string {
				ch := make(chan map[string]string)
				go func() {
					ch <- map[string]string{"ISIN": "123", "Order ID": "456", "Aantal": "1", "Koers": "1.23", "8": "EUR", "Datum": "01-02-2003", "Tijd": "16:01"}
					ch <- map[string]string{"ISIN": "789", "Order ID": "101", "Aantal": "1", "Koers": "1.23", "8": "EUR", "Datum": "01-02-2003", "Tijd": "16:01"}
					close(ch)
				}()
				return ch
			}(),
			done: make(chan bool),
			callback: func(transaction *api.NewTransaction) {
				fmt.Println("Mock callback called with transaction:", transaction)
			},
			expectedOutputs: []string{
				"Mock callback called with transaction: &{123 DeGiro 2003-02-01 16:01:00 +0100 CET  1 {1.23 EUR} { } 456}",
				"Mock callback called with transaction: &{789 DeGiro 2003-02-01 16:01:00 +0100 CET  1 {1.23 EUR} { } 101}",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout && stderr
			realStdout := os.Stdout
			realStderr := os.Stderr

			defer func() {
				os.Stdout = realStdout
				os.Stderr = realStderr
			}()

			rOut, wOut, _ := os.Pipe()
			rErr, wErr, _ := os.Pipe()
			os.Stdout = wOut
			os.Stderr = wErr

			// Run test
			go HandleTransaction(tt.writerChannel, tt.done, tt.callback)

			// Wait for function to finish
			<-tt.done

			// Capture output
			wOut.Close()
			out, _ := ioutil.ReadAll(rOut)
			wErr.Close()
			err, _ := ioutil.ReadAll(rErr)

			for _, expected := range tt.expectedOutputs {
				if !tt.wantErr && !strings.Contains(string(out), expected) {
					t.Errorf("Expected output %q not found in stdout: %s", expected, out)
				} else if tt.wantErr && !strings.Contains(string(err), expected) {
					t.Errorf("Expected output %q not found in stderr: %s", expected, err)
				}
			}
		})
	}
}

func Test_Println(t *testing.T) {
	cet, _ := time.LoadLocation("CET")
	mockTransaction := &api.NewTransaction{Isin: "123", Broker: "DeGiro", Date: time.Date(2003, 02, 01, 16, 01, 0, 0, cet)}

	// Redirect stdout
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run test
	Println(mockTransaction)

	// Capture output
	w.Close()
	out, _ := ioutil.ReadAll(r)

	expectedOutput := "Transaction: &{123 DeGiro 2003-02-01 16:01:00 +0100 CET  0 { } { } }"
	if !strings.Contains(string(out), expectedOutput) {
		t.Errorf("Expected output %q not found in stdout: %s", expectedOutput, out)
	}
}
