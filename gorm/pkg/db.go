package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func InstanceDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("gorm_test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
