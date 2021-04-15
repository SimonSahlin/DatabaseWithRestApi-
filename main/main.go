package main

import (
	"log"
	"net/http"
	"example.com/restapisql/shipmentservice"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	// Create
	createShipmentHandler := http.HandlerFunc(shipmentservice.CreateShipment)
	router.HandleFunc("/shipments", createShipmentHandler).Methods("POST")
	// Read
	router.HandleFunc("/shipments/{shipmentId}", shipmentservice.GetShipment).Methods("GET")
	// Read-all
	router.HandleFunc("/shipments", shipmentservice.GetShipments).Methods("GET")
	// Initialize db connection
	shipmentservice.InitDB()

    log.Fatal(http.ListenAndServe(":8080", router))
}


