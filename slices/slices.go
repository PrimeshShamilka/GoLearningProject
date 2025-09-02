package slices

import "fmt"

// slices are equal to lists/arrays in other languages
func HandleSlices() {
	// same type
	loons := []string{"one", "two", "three"}
	fmt.Printf("%v, %T", loons, loons)

	fmt.Println(len(loons))

	// 0 indexing
	fmt.Println(loons[0])

	fmt.Println(loons[1:])

	//single value range
	for i := range loons {
		fmt.Println(i) // prints indexes
	}
	// double value range
	for i, val := range loons {
		fmt.Printf("%s, %d", val, i)
	}

	// double value range, but ignore index
	for _, val := range loons {
		fmt.Println(val)
	}

	//append
	loons = append(loons, "four")
}
