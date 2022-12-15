package main

import (
	"fmt"
	"sort"
)

func sorting() {
	strs := []string{"c", "a", "b", "cc"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4, 22}
	sort.Ints(ints)
	fmt.Println("Integers:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)

	fruits := []string{"peach", "zzz", "banana", "aaaaaa", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

type byLength []string

func (a byLength) Len() int           { return len(a) }
func (a byLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
