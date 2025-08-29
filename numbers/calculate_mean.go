package numbers

import "fmt"

func CalMean() {
	// var x int = 1
	// var y int = 2

	// var x float64 = 1
	// var y float64 = 2

	// x:= 1.0 // := is shorthand for declaring and initializing a variable
	// y:= 2.0

	x, y := 1.0, 2.0
	/** NOTES:
	1. go won't allow to declare x as int and y as float64, this mathematical 
	is not allowed

	2. go won't allow to compile if we have unused variables
	**/

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result=%v, type of %T\n", mean, mean)
}