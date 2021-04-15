package validating

import (
    "encoding/json"
    "errors"
    "net/http"

    "example.com/restapisql/model"
    "gopkg.in/go-playground/validator.v9"

)

//Checking if inputed countrycodes are valid or not! BELOW
/*--------------------------------------------------------------------------*/
func CheckCountryCode (SenderCountryCode string, ReceiverCountryCode string) error {
    //I would rather have been mapping this, or use the API github.com/biter777/countries if i where to use country names.
    worlds := []string{"AF","AX","AL","DZ","AS","AD","AO","AI","AQ","AG","AR","AM","AW","AU","AT","AZ","BS","BH","BD","BB","BY","BE","BZ",
                     "BJ","BM","BT","BO","BA","BW","BV","BR","IO","BN","BG","BF","BI","KH","CM","CA","CV","KY","CF","TD","CL","CN","CX",
                     "CC","CO","KM","CG","CD","CK","CR","CI","HR","CU","CY","CZ","DK","DJ","DM","DO","EC","EG","SV","GQ","ER","EE","ET",
                     "FK","FO","FJ","FI","FR","GF","PF","TF","GA","GM","GE","DE","GH","GI","GR","GL","GD","GP","GU","GT","GG","GN","GW",
                     "GY","HT","HM","VA","HN","HK","HU","IS","IN","ID","IR","IQ","IE","IM","IL","IT","JM","JP","JE","JO","KZ","KE","KI",
                     "KR","KW","KG","LA","LV","LB","LS","LR","LY","LI","LT","LU","MO","MK","MG","MW","MY","MV","ML","MT","MH","MQ","MR",
                     "MU","YT","MX","FM","MD","MC","MN","ME","MS","MA","MZ","MM","NA","NR","NP","NL","AN","NC","NZ","NI","NE","NG","NU",
                     "NF","MP","NO","OM","PK","PW","PS","PA","PG","PY","PE","PH","PN","PL","PT","PR","QA","RE","RO","RU","RW","BL","SH",
                     "KN","LC","MF","PM","VC","WS","SM","ST","SA","SN","RS","SC","SL","SG","SK","SI","SB","SO","ZA","GS","ES","LK","SD",
                     "SR","SJ","SZ","SE","CH","SY","TW","TJ","TZ","TH","TL","TG","TK","TO","TT","TN","TR","TM","TC","TV","UG","UA","AE",
                     "GB","US","UM","UY","UZ","VU","VE","VN","VG","VI","WF","EH","YE","ZM","ZW"}

    i := ValidateSenderCountry(worlds, SenderCountryCode)
    k := ValidateReceiverCountry(worlds, ReceiverCountryCode)

        if k == true && i == true {
            return nil
        }
	return errors.New("invalid")
}
//Used in CheckCountryCode func to determine if inputed SENDERcountrycode is true or false
func ValidateSenderCountry(slice []string, val string) bool {
        for _, world := range slice {
            if world == val {
                return true
            }
        }
        return false
}
//Used in CheckCountryCode func to determine if inputed RECEIVERcountrycode is true or false
func ValidateReceiverCountry(slice []string, val string) bool {
        for _, world := range slice {
            if world == val {
                return true
            }
        }
	return false
}

/*---------------------------------------------------------------------*/

//Func to write an error message, not if the error accuse but imported and printing a message from API "gopkg.in/go-playground/validator.v9"
func ErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(httpStatusCode)
    resp := make(map[string]string)
    resp["message"] = message
    jsonResp, _ := json.Marshal(resp)
    w.Write(jsonResp)
}

//Func to check for errors in the struct shipment
func ValidateShipment(shipment model.Shipment) error {
    var validate *validator.Validate
    validate = validator.New()
    err := validate.Struct(shipment)
    if err != nil {
        return err
    }
    return nil
}