package main

import "fmt"

func main() {
	precio := 1000.00
	decuentoPorcentaje := 30.00
	porcentajeCalcular := decuentoPorcentaje / 100.00
	decuentoMonto := precio * porcentajeCalcular
	total := precio - decuentoMonto
	fmt.Println("Precio ", precio)
	fmt.Println("descuento ", decuentoPorcentaje, "%")
	fmt.Println("Total a pagar ", total)
}
