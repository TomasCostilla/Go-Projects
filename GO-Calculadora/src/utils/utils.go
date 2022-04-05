package utils

import "fmt"

func Menu(op int) int {

	fmt.Println("***** Ingrese el tipo de operacion *****")
	fmt.Println("***** 1 - Suma")
	fmt.Println("***** 2 - Resta")
	fmt.Println("***** 3 - Division")
	fmt.Println("***** 4 - Division")
	fmt.Printf("Opcion: ")
	fmt.Scan(&op)
	return op
}

func SaveValue(x int) int {
	var result bool
	var num int
	fmt.Printf("Ingrese un numero: ")
	fmt.Scan(&x)
	result = ValidateNum(x)
	if result {
		num = x
	} else {
		fmt.Println("Ingrese un numero mayor a cero!")
		SaveValue(x)
	}
	return num
}

func ValidateNum(x int) bool {
	var result bool /* its false for default */
	if x > 0 {
		result = true
	}
	return result
}
