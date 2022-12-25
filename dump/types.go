package dump

import "fmt"

/*
number at the end of type = number of bits (1/8 byte)

bool
byte = uint8 - can represent ASCII characters
rune = int32 - This integer value is meant to represent a Unicode Code Point
complex64/128 - both real and imaginary part are float32/64
int(32/64 depends on system), int8/16/32/64 - range -2^(n-1) to 2^(n-1)-1
uint(32/64 depends on system), uint8/16/32/64 - range 0 to 2^n-1
uintptr(32/64 bits depends on system) - an unsigned integer type that is large enough to hold any pointer address
float32/64
string - read only slice of bytes
*/
func types() {
	var min uint8 = 255
	fmt.Println(min)
	max := int16(min)
	fmt.Printf("%T\n", max)

	var value interface{} = "str"
	val := value.(string)
	fmt.Println(val)
	val, ok := value.(string)
	fmt.Println(val, ok)
	val2, ok2 := value.(int)
	fmt.Println(val2, ok2)
	// val2 = value.(int) //panic
	// fmt.Println(val2)
}
