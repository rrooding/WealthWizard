package transaction

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"wealth-wizard/degiro-importer/api"
)

func Test_newTransaction(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]string
		want    *api.NewTransaction
		wantErr bool
	}{
		{"Valid data", map[string]string{"ISIN": "123", "Order ID": "456"}, &api.NewTransaction{Isin: "123", Broker: "DeGiro", BrokerId: "456"}, false},
		{"Valid data without order id", map[string]string{"ISIN": "123"}, &api.NewTransaction{Isin: "123", Broker: "DeGiro"}, false},
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
					ch <- map[string]string{"ISIN": "123", "Order ID": "456"}
					close(ch)
				}()
				return ch
			}(),
			done:            make(chan bool),
			callback:        mockCallback,
			expectedOutputs: []string{"Mock callback called with transaction: &{123 DeGiro 456}"},
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
					ch <- map[string]string{"ISIN": "123", "Order ID": "456"}
					ch <- map[string]string{"ISIN": "789", "Order ID": "101"}
					close(ch)
				}()
				return ch
			}(),
			done: make(chan bool),
			callback: func(transaction *api.NewTransaction) {
				fmt.Println("Mock callback called with transaction:", transaction)
			},
			expectedOutputs: []string{
				"Mock callback called with transaction: &{123 DeGiro 456}",
				"Mock callback called with transaction: &{789 DeGiro 101}",
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
