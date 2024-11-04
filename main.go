package main

import (
	"fmt"

	"goBehaviour/overflow"
)

func uintOverflowExample() {

	fmt.Println("=========================================")
	fmt.Println("Integer uint8 overflow example")
	overflow.Uint8overflow()
	fmt.Println("=========================================")

	fmt.Println("=========================================")
	fmt.Println("Integer uint16 overflow example")
	overflow.Uint16overflow()
	fmt.Println("=========================================")

	fmt.Println("=========================================")
	fmt.Println("Integer uint32 overflow example")
	overflow.Uint32overflow()
	fmt.Println("=========================================")

	fmt.Println("=========================================")
	fmt.Println("Integer uint64 overflow example")
	overflow.Uint64overflow()
	fmt.Println("=========================================")
}

func uintUnderflowExample() {
	fmt.Println("=========================================")
	fmt.Println("Integer uint8 underflow example")
	overflow.Uint8underflow()
	fmt.Println("=========================================")

	fmt.Println("=========================================")
	fmt.Println("Integer uint16 underflow example")
	overflow.Uint16underflow()
	fmt.Println("=========================================")

	fmt.Println("=========================================")
	fmt.Println("Integer uint32 underflow example")
	overflow.Uint32underflow()
	fmt.Println("=========================================")

	fmt.Println("=========================================")
	fmt.Println("Integer uint64 underflow example")
	overflow.Uint64underflow()
	fmt.Println("=========================================")
}

func bufferOverflowUsingUnsafe() {
	fmt.Println("=========================================")
	fmt.Println("Example of buffer overflow unsafe")
	handlePanic(func() {
		overflow.BufferOverflowUsingUnsafe()
	})
	fmt.Println("=========================================")

	fmt.Println("=========================================")
	fmt.Println("Example of buffer overflow with a static array unsafe")
	handlePanic(func() {
		overflow.BufferOverflowUsingUnsafeStaticArray()
	})
	fmt.Println("=========================================")

	fmt.Println("=========================================")
	fmt.Println("Example of no buffer overflow safe")
	handlePanic(func() {
		overflow.BufferOverflowWhileSafe()
	})
	fmt.Println("=========================================")

	fmt.Println("=========================================")
	fmt.Println("Example of no buffer overflow with a static array safe")
	handlePanic(func() {
		overflow.BufferOverflowWhileSafeStaticArray()
	})
	fmt.Println("=========================================")
}

// handlePanic executes a function and recovers from any panic, displaying the panic message.
func handlePanic(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	f()
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
