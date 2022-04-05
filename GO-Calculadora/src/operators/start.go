package calculadora

import (
	"fmt"

	"calculadora.com/src/utils"
)

func Start() {

	var num1, num2, res, op int

	x := utils.SaveValue(num1)
	y := utils.SaveValue(num2)
	p := utils.Menu(op)

	if p > 0 || p <= 4 {
		res = Calculate(x, y, p)
		fmt.Println("El resultado es: ", res)
	}

}
