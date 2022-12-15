package main

import (
	"bytes"
	"fmt"
	"os"
)

func bufferFunc() {
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Name ")
	strBuffer.WriteString("Surname")
	fmt.Println(&strBuffer)
	fmt.Println(strBuffer.String())
	strBuffer.Reset()

	fmt.Fprintf(&strBuffer, "It is number: %d, This is a string: %v\n", 10, "Bridge")
	strBuffer.WriteString("[DONE]")
	fmt.Println("The string buffer output is", strBuffer.String())

	var byteString bytes.Buffer
	byteString.Write([]byte("Hello \n"))
	fmt.Fprintf(&byteString, "Hello friends how are you\n")
	byteString.WriteTo(os.Stdout)

	var strBytes = []byte("Hello Name, Surname")
	strBytes = bytes.TrimPrefix(strBytes, []byte("Hello "))
	strBytes = bytes.TrimPrefix(strBytes, []byte("something other"))
	fmt.Printf("%s\n", strBytes)
}
