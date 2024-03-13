package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"wealth-wizard/api/models"
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "stocks_",
		},
	})

	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&models.Transaction{})

	return db
}

func CloseDB(db *gorm.DB) error {
	sqlDB, _ := db.DB()

	return sqlDB.Close()
}
