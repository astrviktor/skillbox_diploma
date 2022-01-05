package billing

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func ParseBillingData(line string) (BillingData, bool) {
	if len(line) != 6 {
		return BillingData{}, false
	}

	var bytes []byte
	runes := []rune(line)
	for _, symbol := range runes {
		switch string(symbol) {
		case "0":
			bytes = append(bytes, 0)
		case "1":
			bytes = append(bytes, 1)
		default:
			return BillingData{}, false
		}
	}
	bytes[0], bytes[5] = bytes[5], bytes[0]
	bytes[1], bytes[4] = bytes[4], bytes[1]
	bytes[2], bytes[3] = bytes[3], bytes[2]

	var digit uint8
	for i := 0; i < 6; i++ {
		digit = digit + bytes[i]*uint8(math.Pow(2, float64(i)))
	}

	CreateCustomer := digit&1 == 1
	Purchase := digit>>1&1 == 1
	Payout := digit>>2&1 == 1
	Recurring := digit>>3&1 == 1
	FraudControl := digit>>4&1 == 1
	CheckoutPage := digit>>5&1 == 1

	elem := BillingData{
		CreateCustomer: CreateCustomer,
		Purchase:       Purchase,
		Payout:         Payout,
		Recurring:      Recurring,
		FraudControl:   FraudControl,
		CheckoutPage:   CheckoutPage,
	}

	return elem, true
}

func StatusBilling(csvFile string) (BillingData, bool) {
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err.Error() + `: ` + csvFile)
		return BillingData{}, false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()

		elem, ok := ParseBillingData(line)

		if ok {
			return elem, true
		}
	}

	return BillingData{}, false
}
