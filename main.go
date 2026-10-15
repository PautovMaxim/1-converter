package main

import "fmt"

func main() {
	const usdToEur = 1 * 0.8926
	const usdToRub = 1 * 83.76
	const EurToRub = usdToRub / usdToEur
	fmt.Println(usdToEur,usdToRub,EurToRub)
}