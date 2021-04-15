package shipmentservice

import (
    "testing"
)

func TestCalculatePriceFromWeight (t *testing.T) {
    /*Description of CalculatePriceFromWeight:
    if actual: 0-10 then exp = 100, if actual: 11-25 then exp = 300, if actual: 26-50 then exp = 500, if actual: 51-1000 then exp = 2000*/
    expectation  := 2000.0 //price output, use float value to check output
    actual := CalculatePriceFromWeight(51)//Change to match question asked
    if actual != expectation {
        t.Error("Something went wrong")
    }
}
