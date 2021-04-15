package shipmentservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"example.com/restapisql/model"
	"example.com/restapisql/validating"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//Initializing the database
func InitDB() {

	var err error
	dataSourceName := "username:password@tcp(localhost:3306)/database_name?parseTime=True" //Databaseconnetiong: Change username and password and database to the correct inputs
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	//db.Exec("CREATE DATABASE shipment_db")
	db.Exec("USE shipment_db")

	// Migration to create tables for Shipments
	db.AutoMigrate(&model.Shipment{})
}

//Creating a shipment with the information from the struct and validating so everything is correct
func CreateShipment(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
    if headerContentType != "application/json"{
        validating.ErrorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
    }
	var shipment model.Shipment
	decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&shipment)
    shipment.SenderCountryCode = strings.ToUpper(shipment.SenderCountryCode) //Change sendercode to Uppercase if input was lowercase
    shipment.ReceiverCountryCode = strings.ToUpper(shipment.ReceiverCountryCode) //Change receivercode to  Uppercase if input was lowercase
    if err != nil{
        validating.ErrorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
        return
    }
    err = validating.ValidateShipment(shipment)
    if err != nil {
        validating.ErrorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
        return
    }
    err = validating.CheckCountryCode(shipment.SenderCountryCode, shipment.ReceiverCountryCode)
    if err != nil {
        validating.ErrorResponse(w, "Bad Request, Sender or Receiver-Code "+err.Error(), http.StatusBadRequest)
        return
    } else {
        validating.ErrorResponse(w, "Success", http.StatusOK)

        var price = CalculatePriceFromWeight(shipment.PackageWeight) //Calculate price based on weight
        var multiple = calculateMultipleByCountry(shipment.SenderCountryCode) //Calculate the multiple based on FROM country
        shipment.Price = price * multiple //Adds the multiple to the price

        db.Create(&shipment) //Adds the shipment to database
    }
}

//Func to calculate the price based on the weight of the package
func CalculatePriceFromWeight (PackageWeight float64) float64 {
    if PackageWeight <= 1000 && PackageWeight >= 51{
        return 2000
    }
    if PackageWeight <= 50 && PackageWeight >= 26{
        return 500
    }
    if PackageWeight <=25 && PackageWeight >=11{
        return 300
    }
    if PackageWeight <= 10 {
        return 100
    }
    return 0
}
//Func for calulating the multiple for the price
func calculateMultipleByCountry(SenderCountryCode string) float64 {
    //Not good use, started off with this and when starting to clean up would have been easier to map these. Recommend -> Create a map for nordic and european for faster checks.
    nordics := []string{"SE","NO","DK","FI"}
    europeans := []string{"AL","AD","AT","AZ","BY","BE","BA","BG","HR","CY","CZ","DK","EE","FI","FR","GE","DE","GR","HU","IS","IE","IT","KZ","XK","LV","LI","LT","LU","MK","MT","MD","MC","ME","NL","NO","PL","PT","RO","RU","SM","RS","SK", "SI","ES","SE","CH","TR","UA","BG","VA"}

    i := findSender(nordics, SenderCountryCode)
        if i == true {
            return 1
        }
    k := findSender(europeans, SenderCountryCode)
        if k == true {
            return 1.5
        }
        if i == false && k == false {
            return 2.5
        }
    fmt.Println("Return 0 for multiple")
	return 0
}
//Used in calculateMultiple and are used to return true if the SenderCode is from Nordic or Europe, otherwise false
func findSender(slice []string, val string) bool {
    for _, nordic := range slice {
        if nordic == val {
            return true
        }
    }
    for _, european := range slice {
        if european == val {
            return true
        }
    }
    return false
}

//The func: When doing a "GET"-request, receives all shipments
func GetShipments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shipments []model.Shipment
	db.Find(&shipments)
	json.NewEncoder(w).Encode(shipments)
}

//The func: When doing a "GET"-request, receives one specific shipment by ID
func GetShipment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputShipmentID := params["shipmentId"]

	var shipment model.Shipment
	db.First(&shipment, inputShipmentID)
	json.NewEncoder(w).Encode(shipment)
}
