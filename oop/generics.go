package oop

import "fmt"

type Ordered interface {
	int | float64 | string
}

// this function has a type T with a constraint of Ordered
func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m, nil
}

func main() {
	fmt.Println(min([]int{5, 3, 6, 2, 1}))
	fmt.Println(min([]float64{5.5, 3.3, 6.6, 2.2, 1.1}))
	fmt.Println(min([]string{"banana", "apple", "cherry"}))
}
