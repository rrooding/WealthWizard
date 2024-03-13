package transaction

import (
	"testing"
  "reflect"

	"wealth-wizard/degiro-importer/api"
)

func Test_newTransaction(t *testing.T) {
  tests := [] struct{
    name string
    data map[string]string
    want *api.NewTransaction
    wantErr bool
  }{
    {"Correct data", map[string]string{"ISIN": "123", "Order ID": "456"}, &api.NewTransaction{Isin: "123", Broker: "DeGiro", BrokerId: "456"}, false},
    {"Correct data without order id", map[string]string{"ISIN": "123"}, &api.NewTransaction{Isin: "123", Broker: "DeGiro"}, false},
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
