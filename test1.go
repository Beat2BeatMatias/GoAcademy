package main

import (
	"./domain"
	"fmt"
)

func main() {
	sitio1 := domain.Site{
		Id:"MLA",
		Name: "Argentina",
	}
	sitio2 := domain.Site{
		Name:"Brasil",
	}
	fmt.Println(sitio1)
	fmt.Println(sitio2.Name)

	categoria1 := domain.Category{
		Id:"MLA5453",
		Name: "Accesorios",
		Site:sitio1,
	}
	categoria2 := domain.Category{
		Id:"MLA5453",
		Name: "Accesorios",
		Site:sitio1,
	}
	categoria3 := domain.Category{
		Site:sitio2,
	}

	fmt.Println(categoria1)
	fmt.Println(categoria2.Site.Name)
	if(categoria1 == categoria2){
		fmt.Println("Las categorías son iguales")
	} else {
		fmt.Println("Las categorías no son iguales")
	}
	if(categoria1 == categoria3){
		fmt.Println("Las categorías son iguales")
	} else {
		fmt.Println("Las categorías no son iguales")
	}
	fmt.Println(categoria1.GetIdsConcateneados())
	fmt.Println(categoria1.Name)
	categoria1.CambiarNombre("Otro nombre")
	fmt.Println(categoria1.Name)
	categoria1.CambiarNombreP("Otro nombre")
	fmt.Println(categoria1.Name)
}
