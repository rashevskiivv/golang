package main

import (
	"fmt"
	"sort"
)

func slices() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	/*
		b := make([]string, len(s), (cap(s)+1)*2)
		for i := range s {
			b[i] = s[i]
		}
		s = b
	*/
	// as same as following

	b1 := make([]string, len(s), (cap(s)+1)*2)
	copy(b1, s)
	s = b1

	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	c := append(a, b...) // b is a slice so u should use ... to append slice to another slice
	fmt.Println(c)
	c = append(c, "aaa First value after sort!")
	fmt.Println(c)
	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})
	fmt.Println(c)

	/*
		c := make([]string, len(s), len(s)+1)
		fmt.Println("emp: ", c)
		num := copy(c, s)
		fmt.Println("num: ", num)
		fmt.Println("cpy: ", c)

		l := s[2:5]
		fmt.Println("sl1: ", l)
		l = s[:2]
		fmt.Println("sl2: ", l)
		l = s[4:]
		fmt.Println("sl3: ", l)

		t := []string{"g", "h", "i"}
		fmt.Println("dcl: ", t)

		twoD := make([][]int, 3)
		for i := 0; i < 3; i++ {
			innerLen := i + 1
			twoD[i] = make([]int, innerLen)
			for j := 0; j < innerLen; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)
	*/
	/*
			package main

		import (
			"golang.org/x/tour/pic"
			"fmt"
		)

		func Pic(dx, dy int) [][]uint8 {
			s := make([][]uint8, dy)
			for i := 0; i < dy; i++ {
				s[i] = make([]uint8, dy)
				s[i][0] = uint8(dx)
			}
			fmt.Println(s)
			return s
		}

		func main() {
			pic.Show(Pic)
		}

	*/
}
