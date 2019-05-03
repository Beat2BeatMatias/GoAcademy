package main

import (
	"./errors"
	"fmt"
	"math"
)

func calcularRaizCuadrada (a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("El número debe ser positivo", a)
	}
	return math.Sqrt(a), nil
}

func main() {
	a := -16.0
	r, err := calcularRaizCuadrada(a)
	if err, ok := err.(*errors.MiError); ok {
		fmt.Println(err.IsNegativo())
	} else {
		fmt.Printf("La raiz cuadrada de %v es %v", a, r)
	}

}
