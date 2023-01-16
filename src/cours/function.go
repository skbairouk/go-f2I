package cours

import (
	"errors"
	"fmt"
)

func Division(nb1 int, nb2 int) (*int, error) {

	if nb2 == 0 {
		return nil, errors.New("Division par 0")
	}
	var calcul = nb1 / nb2
	return &calcul, nil
}

func ArgList(items ...string) { /// toujours dernier argument
	for id, item := range items {
		fmt.Println(id, item)
	}

}

func SwitchCase(name string) {
	switch name {
	case "florant":
		fmt.Println("un admin")
	default:
		fmt.Println("un user")

	}
}
