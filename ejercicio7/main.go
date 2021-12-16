package main

import "fmt"

func main() {
	cliente := map[string]int{"edad": 28, "empleados": 1, "antiguedad": 2, "sueldo": 200000}
	if cliente["edad"] > 22 {
		fmt.Println("aprobado por edad")
	} else {
		fmt.Println("negado por edad")
	}
	if cliente["empleados"] == 1 {
		fmt.Println("aprobado por empleados")
	}
	if cliente["antiguedad"] > 1 {
		fmt.Println("aprobado por antiguedad")
	}
	if cliente["sueldo"] > 100000 {
		fmt.Println("aprobado por sueldo")
	}
}
