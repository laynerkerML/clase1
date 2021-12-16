package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "laynerker"
	sliceWord := strings.Split(text, "")
	fmt.Println("cantidad de letras", len(sliceWord))
	for _, word := range sliceWord {
		fmt.Println("letra: ", word)
	}

}
