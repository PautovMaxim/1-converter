package main

import "fmt"

var amount float64
var originalCurrency string
var targetCurrency string
var result float64

func main() {
	choiseOriginalCurrency()
	getUserAmount()
	choiseTargetCurrency()
	result = calculateCurrency(amount, originalCurrency, targetCurrency)
	fmt.Printf("Итоговое значение: %.2f ", result)
}

func choiseOriginalCurrency() {
	for {
		fmt.Println("Выберите исходную валюту (USD, EUR, RUB): ")
		fmt.Scan(&originalCurrency)	
		if originalCurrency == "USD" || originalCurrency == "EUR" || originalCurrency == "RUB" {
			break;
		}
	}
}

func getUserAmount() {
	for {
		fmt.Println("Введите число: ")
		fmt.Scan(&amount)
		if amount > 0 {
			break;
		}
	}
	
}

func choiseTargetCurrency() {
	for {
		fmt.Println("Выберите целевую валюту (USD, EUR, RUB): ")
		fmt.Scan(&targetCurrency)	
		if originalCurrency == "USD" || originalCurrency == "EUR" || originalCurrency == "RUB" {
			break;
		}
	}
}

func calculateCurrency(amount float64, originalCurrency string, targetCurrency string) float64 {
	if originalCurrency == "USD" && targetCurrency == "EUR" {
		result = amount * 0.8926
	} else if originalCurrency == "EUR" && targetCurrency == "USD" {
		result = amount / 0.8926
	} else if originalCurrency == "USD" && targetCurrency == "RUB" {
		result = amount * 83.76
	} else if originalCurrency == "RUB" && targetCurrency == "USD" {
		result = amount / 83.76
	} else if originalCurrency == "EUR" && targetCurrency == "RUB" {
		result = amount * 92.33
	} else if originalCurrency == "RUB" && targetCurrency == "EUR" {
		result = amount / 92.33
	} 
	return result
}
