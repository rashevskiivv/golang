package dump

import (
	"fmt"
	"os"
)

/*
The behavior of defer statements is straightforward and predictable. There are three simple rules:
1. A deferred function’s arguments are evaluated when the defer statement is evaluated.
2. Deferred function calls are executed in Last In First Out (LIFO) order after the surrounding function returns.
3. Deferred functions may read and assign to the returning function’s named return values.
*/

func deferFunc() {
	f := createFile("data.txt")
	defer closeFile(f)
	writeFile(f)

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	if err := f.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
