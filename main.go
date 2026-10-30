package main

import "fmt"

func main() {
	const usdToEur = 1 * 0.8926
	const usdToRub = 1 * 83.76
	const EurToRub = usdToRub / usdToEur
	fmt.Println(usdToEur,usdToRub,EurToRub)
}

func getUserInput() float64 {
	var userInput float64
	fmt.Println("Введите число: ")
	fmt.Scan(&userInput)
	return userInput
}

func calculateCurrency(amount float64, originalCurrency string, targetCurrency string) {

}