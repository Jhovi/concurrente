package main

import (
	"fmt"
	"log"
	"net/http"

	"./config/server"
)

var PUERTO string

func main() {
	fmt.Print("Ingrese su Puerto: ")
	fmt.Scanf("%s\n", &PUERTO)

	api := server.InitServer(PUERTO)
	log.Fatal(http.ListenAndServe(":"+PUERTO, api.Router()))
}
