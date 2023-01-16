package main

import (
	"fmt"

	"github.com/skbairouk/go-f2i/src/cours"
)

func main() {
	// fmt.Println("hello world")
	// fmt.Println(cours.LeNomDeMaVariablePublic)

	result, err := cours.Division(10, 5)
	fmt.Println(*result, err)

	cours.ArgList("toto", "tata", "titi")

	cours.SwitchCase("florant")

}
