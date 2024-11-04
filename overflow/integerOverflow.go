package overflow

import (
	"fmt"
	"reflect"
)

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
