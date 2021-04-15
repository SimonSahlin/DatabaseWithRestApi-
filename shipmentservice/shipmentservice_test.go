package shipmentservice

import (
    "testing"
)

func TestCalculatePriceFromWeight (t *testing.T) {
    expectation  := CalculatePriceFromWeight(25) //Change to match question asked.
    actual := CalculatePriceFromWeight(11)//Change to match question asked
    if actual != expectation {
        t.Error("Something went wrong")
    }
}
