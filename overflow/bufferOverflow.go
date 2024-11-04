package overflow

import (
	"fmt"
	"unsafe"
)

func BufferOverflowUsingUnsafe() {
	data := []int{1, 2, 3, 4, 5}
	ptr := unsafe.Pointer(&data[0])

	// Attempt to modify data beyond its bounds
	for i := 0; i < 10; i++ {
		element := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*unsafe.Sizeof(data[0])))
		fmt.Println("Original element value", *element)
		*element = i * 10
		fmt.Println("Modified element value", *element)
	}

	fmt.Println(data) // Undefined behavior and may lead to data corruption
}

func BufferOverflowUsingUnsafeStaticArray() {
	data := [5]int{1, 2, 3, 4, 5}
	ptr := unsafe.Pointer(&data[0])

	// Attempt to modify data beyond its bounds
	for i := 0; i < 10; i++ {
		element := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*unsafe.Sizeof(data[0])))
		fmt.Println("Original element value", *element)
		*element = i * 10
		fmt.Println("Modified element value", *element)
	}

	fmt.Println(data) // Undefined behavior and may lead to data corruption
}

func BufferOverflowWhileSafe() {
	data := []int{1, 2, 3, 4, 5}

	// Attempt to modify data beyond its bounds
	for i := 0; i < 10; i++ {
		fmt.Println("Original element value", data[i])
		data[i] = i * 10
		fmt.Println("Modified element value", data[i])
	}

	fmt.Println(data) // Undefined behavior and may lead to panic
}

func BufferOverflowWhileSafeStaticArray() {
	data := [5]int{1, 2, 3, 4, 5}

	// Attempt to modify data beyond its bounds
	for i := 0; i < 10; i++ {
		fmt.Println("Original element value", data[i])
		data[i] = i * 10
		fmt.Println("Modified element value", data[i])
	}

	fmt.Println(data) // Undefined behavior and may lead to panic
}
