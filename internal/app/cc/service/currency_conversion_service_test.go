package service

import "testing"

func TestConvertCurrency_Success(t *testing.T) {
	// Initialize conversion rates for testing
	ConversionRates = map[string]float64{
		"USD": 1.0,
		"EUR": 0.85,
		"INR": 74.0,
	}

	tests := []struct {
		name             string
		senderCurrency   string
		receiverCurrency string
		amount           float64
		expectedAmount   float64
	}{
		{"Valid conversion USD to EUR", "USD", "EUR", 100, 85.00},
		{"Valid conversion EUR to INR", "EUR", "INR", 100, 8705.88},
		{"No conversion needed", "USD", "USD", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, err := ConvertCurrency(tt.senderCurrency, tt.receiverCurrency, tt.amount)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if amount != tt.expectedAmount {
				t.Errorf("expected amount: %v, got: %v", tt.expectedAmount, amount)
			}
		})
	}
}

func TestConvertCurrency_Failure(t *testing.T) {
	// Initialize conversion rates for testing
	ConversionRates = map[string]float64{
		"USD": 1.0,
		"EUR": 0.85,
		"INR": 74.0,
	}

	tests := []struct {
		name             string
		senderCurrency   string
		receiverCurrency string
		amount           float64
		expectError      bool
	}{
		{"Negative Input amount", "USD", "EUR", -100, true},
		{"0 Input amount", "USD", "EUR", 0, true},
		{"Invalid input sender currency", "", "EUR", 100, true},
		{"Invalid input receiver currency", "USD", "", 100, true},
		{"Unsupported sender currency", "JPY", "USD", 100, true},
		{"Unsupported receiver currency", "USD", "JPY", 100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ConvertCurrency(tt.senderCurrency, tt.receiverCurrency, tt.amount)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
		})
	}
}
