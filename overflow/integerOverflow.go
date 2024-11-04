package overflow

import (
	"fmt"
	"reflect"
)

func Uint8overflow() {
	var i uint8 = 255
	fmt.Println("Type of i is: ", reflect.TypeOf(i))
	fmt.Println("Value of i + 0 is ", i)
	i += 1
	fmt.Println("Value of i + 1 is ", i)
	i += 1
	fmt.Println("Value of i + 2 is ", i)
	i += 1
	fmt.Println("Value of i + 3 is ", i)

	fmt.Println("Type of i is: ", reflect.TypeOf(i))
}

func Uint16overflow() {
	var i uint16 = 65535
	fmt.Println("Type of i is: ", reflect.TypeOf(i))
	fmt.Println("Value of i + 0 is ", i)
	i += 1
	fmt.Println("Value of i + 1 is ", i)
	i += 1
	fmt.Println("Value of i + 2 is ", i)
	i += 1
	fmt.Println("Value of i + 3 is ", i)

	fmt.Println("Type of i is: ", reflect.TypeOf(i))
}

func Uint32overflow() {
	var i uint32 = 4294967295
	fmt.Println("Type of i is: ", reflect.TypeOf(i))
	fmt.Println("Value of i + 0 is ", i)
	i += 1
	fmt.Println("Value of i + 1 is ", i)
	i += 1
	fmt.Println("Value of i + 2 is ", i)
	i += 1
	fmt.Println("Value of i + 3 is ", i)

	fmt.Println("Type of i is: ", reflect.TypeOf(i))
}

func Uint64overflow() {
	var i uint64 = 18446744073709551615
	fmt.Println("Type of i is: ", reflect.TypeOf(i))
	fmt.Println("Value of i + 0 is ", i)
	i += 1
	fmt.Println("Value of i + 1 is ", i)
	i += 1
	fmt.Println("Value of i + 2 is ", i)
	i += 1
	fmt.Println("Value of i + 3 is ", i)

	fmt.Println("Type of i is: ", reflect.TypeOf(i))
}

func Uint8underflow() {
	var i uint8 = 0
	fmt.Println("Type of i is: ", reflect.TypeOf(i))
	fmt.Println("Value of i - 0 is ", i)
	i -= 1
	fmt.Println("Value of i - 1 is ", i)
	i -= 1
	fmt.Println("Value of i - 2 is ", i)
	i -= 1
	fmt.Println("Value of i - 3 is ", i)

	fmt.Println("Type of i is: ", reflect.TypeOf(i))
}

func Uint16underflow() {
	var i uint16 = 0
	fmt.Println("Type of i is: ", reflect.TypeOf(i))
	fmt.Println("Value of i - 0 is ", i)
	i -= 1
	fmt.Println("Value of i - 1 is ", i)
	i -= 1
	fmt.Println("Value of i - 2 is ", i)
	i -= 1
	fmt.Println("Value of i - 3 is ", i)

	fmt.Println("Type of i is: ", reflect.TypeOf(i))
}

func Uint32underflow() {
	var i uint32 = 0
	fmt.Println("Type of i is: ", reflect.TypeOf(i))
	fmt.Println("Value of i - 0 is ", i)
	i -= 1
	fmt.Println("Value of i - 1 is ", i)
	i -= 1
	fmt.Println("Value of i - 2 is ", i)
	i -= 1
	fmt.Println("Value of i - 3 is ", i)

	fmt.Println("Type of i is: ", reflect.TypeOf(i))
}

func Uint64underflow() {
	var i uint64 = 0
	fmt.Println("Type of i is: ", reflect.TypeOf(i))
	fmt.Println("Value of i - 0 is ", i)
	i -= 1
	fmt.Println("Value of i - 1 is ", i)
	i -= 1
	fmt.Println("Value of i - 2 is ", i)
	i -= 1
	fmt.Println("Value of i - 3 is ", i)

	fmt.Println("Type of i is: ", reflect.TypeOf(i))
}
