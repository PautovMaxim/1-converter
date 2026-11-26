package main

import "fmt"

func main() {
	userChoiseOperation := choiseOfOperation() //запрашивает ввод пользователя и возвращает наименование оперции
	userNumbers := getUserNumbers() //запрашивает числа пользователя и возвращает массив чисел

	fmt.Println(userChoiseOperation)
	fmt.Println(userNumbers)
}

func choiseOfOperation() string { //запрашивает ввод пользователя и возвращает наименование оперции
	var userChoise string
	for {
		fmt.Println("Выберите операцию(AVG, SUM, MED): ")
		fmt.Scan(&userChoise)
		if userChoise == "AVG" || userChoise == "SUM" || userChoise == "MED" {
			break;
		}
		fmt.Println("Некорректно указана операция") 
	}
	return userChoise
}

func getUserNumbers() []float64 {
	userNumbers := []float64{}
	for {
		var userNumber float64 
		fmt.Println("Введите число(для выхода нажмите - n): ")
		_, err := fmt.Scan(&userNumber)
		if err != nil {
			break;
		}
		userNumbers = append(userNumbers, userNumber)
	}
	return userNumbers
}
