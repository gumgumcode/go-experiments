package main

func main() {

	num := 5
	modifyValue(num)
	// `num` remains 5, as the function received a copy of the value

	data := []int{1, 2, 3}
	modifySlice(data)
	// `data` is now []int{99, 2, 3}

	num2 := 5
	modifyWithPointer(&num2)
	// `num2` is now 10, as the function received a reference to `num2` via a pointer

}

// [1] Passing by Value:
// In Go, all function arguments are passed by value by
// default. When you pass a variable as an argument to a function, a copy of
// the value (not a reference to the original) is made, and this copy is
// used within the function. Changes made to the parameter within the
// function do not affect the original variable outside the function. This
// is true for most basic types like integers, strings, and structs.

func modifyValue(x int) {
	x = x * 2
}

// [2] Passing by Reference:
// Go does not support traditional "pass by reference"
// like some other languages. However, when you pass a pointer to an object
// (like a slice, map, or a custom struct) as an argument to a function, you
// are effectively passing a reference to that object. Changes made to the
// object through the pointer will affect the original object outside the
// function.

func modifySlice(slice []int) {
	slice[0] = 99
}

// [3] Passing by Pointer:
// When you pass a pointer to a variable as an
// argument to a function, you are passing a reference to that variable.
// This means that changes made to the variable through the pointer will
// affect the original variable outside the function.

func modifyWithPointer(x *int) {
	*x = *x * 2
}
