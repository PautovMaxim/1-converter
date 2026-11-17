package main

import "fmt"

var amount float64
var originalCurrency string
var targetCurrency string
var result float64

func main() {
	originalCurrency = readCurrency ("Выберите целевую валюту (USD, EUR, RUB): ")
	getUserAmount()
	targetCurrency = readCurrency ("Выберите целевую валюту (USD, EUR, RUB): ")
	result = calculateCurrency(amount, originalCurrency, targetCurrency)
	fmt.Printf("Итоговое значение: %.2f ", result)
}

func readCurrency (prompt string) string {
	var currency string
	for {
		fmt.Println(prompt)
		fmt.Scan(&currency)
		if currency == "USD" || currency == "EUR" || currency == "RUB" {
			return currency
		}

		fmt.Println("Некорректная валюта. Попробуйте ещё раз")
	}
}

func getUserAmount() {
	for {
		fmt.Println("Введите число: ")
		_, err := fmt.Scan(&amount)
		
		if err != nil {
			fmt.Println("Ошибка! Введите корректное число.")
			continue
		}

		if amount > 0 {
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
	} else if originalCurrency == targetCurrency {
		result = amount
	}
	return result
}
