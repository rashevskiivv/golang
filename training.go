package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/julienschmidt/httprouter"
)

func train() {
	log.Println("read config")
	var cfg map[string]string
	cleanenv.ReadConfig("config.yml", &cfg)

	router := httprouter.New()
	router.GET("/", IndexHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg["port"]), router))
}

func IndexHandler(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	w.WriteHeader(200)
	w.Write([]byte("Hello From VSCode!"))
}
