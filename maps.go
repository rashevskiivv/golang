package main

import "fmt"

func maps() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "app": 0, "bar": 2} //map sort by keys
	fmt.Println("map:", n)
	/*
		package main

	import (
		"golang.org/x/tour/wc"
		"strings"
	)

	func WordCount(s string) map[string]int {
		m := map[string]int{}
		words := strings.Fields(s)
		for _, word := range words {
			if _, ok := m[word]; ok {
				m[word] += 1
			} else {
				m[word] = 1
			}
		}
		return m
	}

	func main() {
		wc.Test(WordCount)
	}

	*/
}
