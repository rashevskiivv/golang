package dump

import (
	"fmt"
	"math"
)

const str string = "constant"

const ( //enum
	Black = iota
	Grey
	White
)

func constants() {
	fmt.Println(str)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

const (
	one = iota*2 + 1 // укажите здесь формулу с iota
	three
	five
	seven
	nine
	eleven
)

//func main() {
//	fmt.Println(one, three, five, seven, nine, eleven)
//}
