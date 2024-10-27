package service

import (
	"fmt"
	"strconv"
)

var ConversionRates map[string]float64

func ConvertCurrency(senderCurrency, receiverCurrency string, amount float64) (float64, error) {
	if amount <= 0 || senderCurrency == "" || receiverCurrency == "" {
		return -1, fmt.Errorf("invalid input")
	}
	if senderCurrency == receiverCurrency { // No conversion needed
		return amount, nil
	}

	senderRate, senderOk := ConversionRates[senderCurrency]
	receiverRate, receiverOk := ConversionRates[receiverCurrency]

	if !senderOk {
		return 0, fmt.Errorf("unsupported sender currency: %s", senderCurrency)
	}
	if !receiverOk {
		return 0, fmt.Errorf("unsupported receiver currency: %s", receiverCurrency)
	}

	// Format the converted amount to two decimal places
	convertedAmount, err := strconv.ParseFloat(fmt.Sprintf("%.2f", (amount/senderRate)*receiverRate), 64)
	if err != nil {
		return 0, fmt.Errorf("failed to format converted amount: %v", err)
	}

	return convertedAmount, nil
}
