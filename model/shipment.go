package model


type Shipment struct {
	ShipmentID          uint      `json:"shipmentId" gorm:"primary_key"`
	SenderName          string    `json:"senderName" validate:"required,lte=30"`
	SenderEmail         string    `json:"senderEmail" validate:"required,email"`
	SenderAddress       string    `json:"senderAddress" validate:"required,lte=100"`
	SenderCountryCode   string    `json:"senderCountrycode" validate:"required,len=2,alpha"` //Would rather build with CountryName, to use github.com/biter777/countries API
    ReceiverName        string    `json:"receiverName" validate:"required,lte=30"`
    ReceiverEmail       string    `json:"receiverEmail" validate:"required,email"`
    ReceiverAddress     string    `json:"receiverAddress" validate:"required,lte=100"`
    ReceiverCountryCode string    `json:"receiverCountrycode" validate:"required,len=2,alpha"`
    PackageWeight       float64   `json:"packageWeight" validate:"required,numeric,lte=1000"`
    Price               float64   `json:"price"`
}