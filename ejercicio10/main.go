package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	for k, employee := range employees {
		if k == "Nahuel" {
			fmt.Println("nombre:", k, " Edad:", employee)
		}
		if employee >= 21 {
			fmt.Println("mayor de 21 nombre: ", k, " Edad:", employee)
		}
	}

	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("/++++++++++++++++/")

	for k, employee := range employees {
		if k == "Nahuel" {
			fmt.Println("nombre:", k, " Edad:", employee)
		}
		if employee >= 21 {
			fmt.Println("mayor de 21 nombre: ", k, " Edad:", employee)
		}
	}

}
