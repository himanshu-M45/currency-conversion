package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Currency struct {
	Code      string  `xml:"code"`
	BaseValue float64 `xml:"baseValue"`
}

type ConversionRates struct {
	Currencies []Currency `xml:"currency"`
}

func LoadConversionRates(filePath string) (map[string]float64, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open XML file: %v", err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var conversionRates ConversionRates
	err = xml.Unmarshal(byteValue, &conversionRates)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %v", err)
	}

	rates := make(map[string]float64)
	for _, currency := range conversionRates.Currencies {
		rates[currency.Code] = currency.BaseValue
	}

	return rates, nil
}
