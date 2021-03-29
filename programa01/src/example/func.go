package main

import (
	"errors"
	"math"
)

func main() {
	resultadoSoma := sum(2, 4)
	println(resultadoSoma)

	resultadoSqlt, err := sqlt(-16)

	if err != nil {
		println(err)
	} else {
		println(resultadoSqlt)
	}

}

func sum(a int, b int) int {
	return a + b
}

func sqlt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Numero negativo")
	}
	return math.Sqrt(x), nil
}
