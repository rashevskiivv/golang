package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db gorm.DB

func InstanceDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	db1, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db1.AutoMigrate(Product{})
	if err != nil {
		log.Fatal(err)
	}
	db = *db1
}

func Create(productToInsert Product) (uint, error) {
	dbAffected := db.Create(&productToInsert)
	_, err := checkAffectedDB(dbAffected)
	if err != nil {
		return 0, err
	}
	return productToInsert.ID, nil
}

func ReadAll() ([]Product, error) {
	var products []Product
	if _, err := checkAffectedDB(db.Find(&products)); err != nil {
		return products, err
	}
	return products, nil
}

func Read(where string) (Product, error) {
	var product Product
	dbAffected := db.Where(where).First(&product)
	if err := dbAffected.Error; err != nil {
		return Product{}, err
	}
	if dbAffected.RowsAffected > 0 {
		return product, nil
	}
	return Product{}, errors.New("Read func has affected 0 rows with condition: " + where)
}

func Delete(where string) (bool, error) {
	var product Product
	return checkAffectedDB(db.Where(where).Delete(&product))
}

func DeleteAll() (bool, error) {
	return checkAffectedDB(db.Exec("DELETE FROM products"))
}

func Update(id uint, productToUpdate Product) (Product, error) {
	var product Product
	dbAffected := db.Model(&product).Where(fmt.Sprintf("id=%v", id)).Updates(productToUpdate)
	_, err := checkAffectedDB(dbAffected)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func checkAffectedDB(tx *gorm.DB) (bool, error) {
	if err := tx.Error; err != nil {
		return false, err
	}
	if tx.RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}

func init() {
	InstanceDB()
}
