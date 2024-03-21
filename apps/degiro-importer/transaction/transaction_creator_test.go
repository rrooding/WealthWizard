package transaction

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"wealth-wizard/degiro-importer/api"
)

// Mock client for CreatorFunc
type MockClient struct{}

func (m *MockClient) CreateTransaction(input api.NewTransaction) (*api.CreateTransactionResponse, error) {
	fmt.Println("Mock client called with transaction:", input)

	return &api.CreateTransactionResponse{
		CreateTransaction: api.CreateTransactionCreateTransaction{Id: 1},
	}, nil
}

func Test_CreatorFunc(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Amsterdam")
	date := time.Date(2003, 02, 01, 16, 01, 0, 0, loc)

	mockClient := &MockClient{}

	// Redirect stdout
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Create the function
	creatorFunc := CreatorFunc(mockClient)
	if creatorFunc == nil {
		t.Errorf("CreatorFunc() = nil, want a function")
	}

	// Check if the function is working
	mockTransaction := &api.NewTransaction{Isin: "123", Broker: "DeGiro", BrokerId: "456", Amount: 1, Date: date}
	creatorFunc(mockTransaction)

	// It calls the mock client
	w.Close()
	out, _ := ioutil.ReadAll(r)

	expectedOutput := "Mock client called with transaction: {123 DeGiro 2003-02-01 16:01:00 +0100 CET  1 { } { } 456}"
	if !strings.Contains(string(out), expectedOutput) {
		t.Errorf("Expected output %q not found in stdout: %s", expectedOutput, out)
	}
}
