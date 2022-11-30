package main

import (
	"fmt"
	"strconv"
)

func for_loop() {
	i := 0

	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	i = 0
	for {
		fmt.Println("loop " + strconv.Itoa(i))
		if i == 4 {
			break
		} else {
			i++
		}
	}

	for n := 0; n != 4; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
