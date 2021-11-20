package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=realm password=utongnimacmac dbname=ndfy_development port=5432 sslmode=disable TimeZOne=ASia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
