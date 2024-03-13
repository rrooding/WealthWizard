package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model

	ISIN     string
	Broker   string
	BrokerID string
}
