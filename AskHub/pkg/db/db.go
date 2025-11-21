package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(connStr string) *Db{
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	return &Db{db}
}