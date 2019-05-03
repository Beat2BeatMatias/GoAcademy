package main

import (
	"fmt"
)

func funcion1(){
	fmt.Println("Imprime función 1")
}
func funcion2(){
	fmt.Println("Imprime función 2")
}
func funcion3(){
	fmt.Println("Imprime función 3")
}

func main() {
	defer funcion1()
	defer funcion2()
	for i := 0; i<100; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")
	funcion3()
}