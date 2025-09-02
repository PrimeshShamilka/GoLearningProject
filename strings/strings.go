package strings

import "fmt"

func HandleStrings() {
	book := "Test book"
	fmt.Println(book)
	fmt.Println(len(book))

	// strings in go are immutable
	// book[0] = "1"

	// slicing
	fmt.Println(book[4:8])
	fmt.Println(book[:4])
	fmt.Println(book[1:])

	// concatenation
	fmt.Println("t" + book[1:])

	// multi-line
	poem := `
		This is a 
		test multi
		line string
		...
		`
	fmt.Println(poem)

	// convert int to str
	n := 58
	s := fmt.Sprintf("%d", n)
	fmt.Printf("%q %T", s, s)
}
