package calculadora

func Calculate(x int, y int, op int) int {
	var result int
	switch op {
	case 1:
		result = Suma(x, y)
	case 2:
		result = Resta(x, y)
	case 3:
		result = Division(x, y)
	case 4:
		result = Division(x, y)
	default:
		result = 0
	}

	return result
}

func Suma(x int, y int) int {
	result := x + y
	return result
}

func Resta(x int, y int) int {
	result := x - y
	return result
}

func Multiplicacion(x int, y int) int {
	result := x * y
	return result
}

func Division(x int, y int) int {
	result := x / y
	return result
}
