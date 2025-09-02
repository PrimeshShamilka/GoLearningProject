package conditionals

import "fmt"

func CheckConditions() {
	x := 10

	if x > 5 {
		fmt.Println("X is big")
	}

	if x > 100 {
		fmt.Println("X is very big")
	} else {
		fmt.Println("X is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("X is right")
	}

	if x < 20 || x > 30 {
		fmt.Println("X is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is half of b")
	}

	// switch-case
	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	switch {
	case n > 100:
		fmt.Println("n is very big")
	case n > 10:
		fmt.Println("n is big")
	default:
		fmt.Println("other")
	}

}
