package dump

import (
	"fmt"
	"golang/mymath"
)

func packages() {
	values := []float64{0, 1, 2, 3, 4, 5, 6, 13}
	avg := mymath.Avg(values)
	fmt.Println(avg)
}
