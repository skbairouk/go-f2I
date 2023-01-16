package cours

import "fmt"

func If(nb int) {
	if nb == 0 {
		fmt.Println(nb)
	}
}

func Ifelse(nb int) bool {
	if nb == 0 {
		return true
	}
	return false
}
