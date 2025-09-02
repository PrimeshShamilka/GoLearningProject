package functions

import (
	"fmt"
	"math"
)

func main() {
	val := 10
	double(val)     // won't reflect the change
	doublePtr(&val) // will reflect the change
}

func Add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

// slices are passed as pointers
func doubleAt(values []int, i int) {
	values[i] *= 2
}

// when a is passed from somewhere else, it makes a copy
// so this method won't double the class parameter
func double(a int) {
	a *= 2
}

// using pointers will update the class parameter
func doublePtr(a *int) {
	*a *= 2
}

// Errors
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.00, fmt.Errorf("negative")
	}
	return math.Sqrt(n), nil
}

// Defer
func worker() {
	r1, err := acquire("A")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer release(r1) // this will run at the end of this method, i.e. after printing 'worker'
	fmt.Println("worker")
}

func acquire(name string) (string, error) {
	return name, nil
}

// Ex - cleaning other resources like ports, etc
func release(name string) {
	fmt.Printf("Cleaning up %s", name)
}
