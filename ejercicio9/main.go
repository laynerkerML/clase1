package main

import "fmt"

func main() {
	alumnos := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	for _, alumn := range alumnos {
		fmt.Println("alumno: ", alumn)
	}
	fmt.Println("/++++++++/")
	alumnos = append(alumnos, "Gabriela")

	for _, alumn := range alumnos {
		fmt.Println("alumno: ", alumn)
	}
}
