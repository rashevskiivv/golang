package main

import "os"

func panicFunc() {
	_, err := os.Create("file||\\//")
	if err != nil {
		panic(err)
	}
}
