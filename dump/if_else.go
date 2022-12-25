package dump

import (
	"fmt"
	"math/rand"
)

func if_else() {
	maxRand := 100
	for i := 0; i < 10; i++ {
		if num := rand.Intn(maxRand) - 25; num < 0 {
			fmt.Println(num, " is negative")
		} else if num < 10 {
			fmt.Println(num, "has 1 digit")
		} else {
			fmt.Println(num, "has multiple digits")
		}
	}
}
