package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model

	Isin            string
	Broker          string
	BrokerID        string
	Date            time.Time
	Exchange        string
	Amount          int
	Price           decimal.Decimal `gorm:"type:decimal(20,8);"`
	Currency        string
	TransactionCost *decimal.Decimal `gorm:"type:decimal(20,8);"`
}

func GenerateBrokerId(t *Transaction) (string, error) {
	data := fmt.Sprintf("%s%s%v%s%d", t.Isin, t.Broker, t.Date, t.Exchange, t.Amount)
	id := uuid.NewMD5(uuid.NameSpaceOID, []byte(data))

	return id.String(), nil
}
