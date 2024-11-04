package main

import (
	"fmt"
	"reflect"

	"goBehaviour/overflow"
)

// handlePanic executes a function and recovers from any panic, displaying the panic message.
func handlePanic(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	f()
}

// func outputFunc(msg string, f func()) {
// 	fmt.Println("=========================================")
// 	fmt.Println(msg)
// 	handlePanic(func() {
// 		f()
// 	})
// 	fmt.Println("=========================================")
// }

func outputFunc(msg string, f interface{}, args ...interface{}) {
	fmt.Println("=========================================")
	fmt.Println(msg)
	handlePanic(func() {
		fn := reflect.ValueOf(f)               // converts the function f from an interface{} type to a reflect.Value type for dynamic invocation of the function
		in := make([]reflect.Value, len(args)) // A slice in of reflect.Value is created to hold the converted arguments

		// This conversion is necessary because the Call method of reflect.Value expects its arguments in reflect.Value form.
		for i, arg := range args { // A slice in of reflect.Value is created to hold the converted arguments.
			in[i] = reflect.ValueOf(arg) // each argument is converted to a reflect.Value and stored in the in slice
		}

		fn.Call(in) // invokes the function f using the Call method of the reflect.Value type, passing the in slice of arguments.
	})
	fmt.Println("=========================================")
}

func uintOverflowExample() {

	outputFunc("Integer uint8 overflow example", overflow.Uint8overflow)
	outputFunc("Integer uint16 overflow example", overflow.Uint16overflow)
	outputFunc("Integer uint32 overflow example", overflow.Uint32overflow)
	outputFunc("Integer uint64 overflow example", overflow.Uint64overflow)

}

func uintUnderflowExample() {

	outputFunc("Integer uint8 underflow example", overflow.Uint8underflow)
	outputFunc("Integer uint16 underflow example", overflow.Uint16underflow)
	outputFunc("Integer uint32 underflow example", overflow.Uint32underflow)
	outputFunc("Integer uint64 underflow example", overflow.Uint64underflow)
}

func bufferOverflowUsingUnsafe() {

	outputFunc("Example of buffer overflow unsafe", overflow.BufferOverflowUsingUnsafe)
	outputFunc("Example of buffer overflow with a static array unsafe", overflow.BufferOverflowUsingUnsafeStaticArray)
	outputFunc("Example of no buffer overflow safe", overflow.BufferOverflowWhileSafe)
	outputFunc("Example of no buffer overflow with a static array safe", overflow.BufferOverflowWhileSafeStaticArray)

}

func main() {

	var userChoice string
	fmt.Println("Press 1 to see uintOverflowExample")
	fmt.Println("Press 2 to see uintUnderflowExample")
	fmt.Println("Press 3 to see bufferOverflowUsingUnsafe")
	fmt.Println("Enter your choice")
	fmt.Scanln(&userChoice)

	switch userChoice {
	case "1":
		uintOverflowExample()
	case "2":
		uintUnderflowExample()
	case "3":
		bufferOverflowUsingUnsafe()
	default:
		fmt.Println("Invalid input")
	}

}
