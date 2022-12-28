package main

import (
	"crypto/rand"
	database "github.com/rashviach/golang/gorm/internal"
	"log"
	"math/big"
)

func main() {
	database.InstanceDB()

	nBig, err := rand.Int(rand.Reader, big.NewInt(200))
	if err != nil {
		panic(err)
	}
	price := nBig.Int64()
	idx, err := database.Create(database.Product{
		Code:  string(rune(price)),
		Price: uint(price),
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Entry created with id %v\n", idx)

	allProducts, err := database.ReadAll()
	if err != nil {
		panic(err)
	}
	for idx, product := range allProducts {
		log.Printf("%v: %v %v %v\n", idx, product.ID, product.Price, product.Code)
	}

	product, err := database.Read("id=2")
	if err != nil {
		panic(err)
	}
	log.Printf("%v %v %v", product.ID, product.Price, product.Code)

	updatedProduct, err := database.Update(8, database.Product{
		Code:  "Z13",
		Price: 1488,
	})
	if err != nil {
		panic(err)
	}
	log.Println(updatedProduct)

	isDeleted, err := database.Delete("id=20")
	if err != nil {
		panic(err)
	}
	log.Println(isDeleted)

	isDeleted, err = database.DeleteAll()
	if err != nil {
		panic(err)
	}
	log.Println(isDeleted)
}
