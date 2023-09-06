package main

import (
	"fmt"
	"os"
)

func main() {

	// Reading from a file.
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	fmt.Println("File content:")
	fmt.Println(string(content))

	// Writing to a file.
	data := []byte("Hello world!")
	err = os.WriteFile("output.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}
	fmt.Println("Data written to the file successfully.")

	// Make a directory.
	err = os.Mkdir("my_dir", 0755)
	if err != nil {
		fmt.Println("Error creating directory.")
		return
	}

	// Delete a file.
	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error deleting the file.")
	}

	// Delete a directory.
	err = os.Remove("my_dir")
	if err != nil {
		fmt.Println("Error deleting the directory.")
	}

	// Check if file exists.
	if _, err = os.Stat("example.txt"); os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		fmt.Println("File exists.")
	}
}
