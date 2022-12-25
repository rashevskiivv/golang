package dump

import "fmt"

func functions() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	a, b := vals()
	_, c := vals()
	fmt.Println(a, b, c)

	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	nextInt = intSeq()
	fmt.Println(nextInt())

	fmt.Println(fact(7))
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))

	fmt.Println(namedReturnValues())
}

func plus(a, b int) int {
	return a + b
}

func vals() (int, int) {
	return 3, 7
}

// variadic functions
func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// closures
// A closure is a function value that references variables from outside its body.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func namedReturnValues() (x, y int) {
	x = 10
	y = 11
	return
}
