package oop

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance float64
	Expires time.Time
}

func main() {
	b1 := Budget{"Food", 245.22, time.Now().Add(7*time.Hour)}
	fmt.Println(b1)
	// print with more details
	fmt.Printf("%#v", b1)
	fmt.Println(b1.CampaignID)

	b2 := Budget{
		Balance: 100.11,
		CampaignID: "Car",
	}
	fmt.Printf("%#v", b2)

	var b3 Budget
	fmt.Printf("%#v", b3)

	/**
	Access specifiers:
	- Everything that starts with a Capital letter is accessible from outside the package (Exported symbols)
	- Otherwise it's accessible only within the current package (Unexported symbols)
	**/


	// CREATING METHODS ON STRUCTS
	func (b Budget) TimeLeft() time.Duration {
		return b.Expires.Sub(time.Now().UTC())
	}

}	

