package maps

import "fmt"

// key value pair data structure
func HandleMaps() {
	stocks := map[string]float64{
		"AMZN": 1033.00,
		"GOOG": 2353.22,
		"TESL": 33434.22,
	}
	fmt.Println(len(stocks))

	// get a value, get zero is key not found
	fmt.Println(stocks["GOOG"])

	// check if val is in the map or not
	val, ok := stocks["ABCD"]
	if !ok {
		fmt.Println("not in stocks")
	} else {
		fmt.Println(val)
	}

	// set a value
	stocks["ABCD"] = 234324.334

	// print items
	for keys := range stocks {
		fmt.Println(keys)
	}
}
