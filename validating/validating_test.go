package validating

import (
    "testing"
)

func TestCheckCountryCode(t *testing.T){
     expectation  := CheckCountryCode("SE", "DK") //Dont change, these are correct
     actual := CheckCountryCode("SE", "US")//Change to se if the CountryCode exists, only UpperLetters
     if actual != expectation {
         t.Error("Does not exist")
     }
 }

