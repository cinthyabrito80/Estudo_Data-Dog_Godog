package main

import (
	"fmt"
)

func main() {

	var numero int = 2500

	if numero%2 == 0 {
		fmt.Println("numero :", numero, "é par")
	} else {
		fmt.Println("numero :", numero, "é impar")
	}

}
