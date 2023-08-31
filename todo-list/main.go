package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tasks := make([]string, 0)

	fmt.Println("Welcome to the command line todo list app!")
	fmt.Print("\n1. Add Task\n2. List Tasks\n3. Delete\n4. Quit\n")

	for {
		fmt.Print("\nEnter an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			task := scanner.Text()
			tasks = append(tasks, task)
			fmt.Println("Task added successfully!")
		case 2:
			if len(tasks) == 0 {
				fmt.Println("No tasks to display.")
			} else {
				printTasks(tasks)
			}
		case 3:
			if len(tasks) == 0 {
				fmt.Println("No tasks to delete.")
			} else {
				var indexToDelete int
				fmt.Print("\nEnter task number to delete: ")
				fmt.Scan(&indexToDelete)
				indexToDelete -= 1
				if indexToDelete >= 0 && indexToDelete < len(tasks) {
					tasks = append(tasks[:indexToDelete], tasks[indexToDelete+1:]...)
				} else {
					fmt.Println("Invalid task number!")
				}
				printTasks(tasks)
			}
		case 4:
			fmt.Println("Closed!")
			return
		default:
			fmt.Println("Invalid choice. Please select again.")
		}
	}
}

func printTasks(tasks []string) {
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func scanLineByLine() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("You entered: ", text)
	}
}

/*
fmt.Scan():
- It doesn't work well for reading lines of text
or unformatted input unless you're explicitly using format specifiers.
- Commonly used for simple console input where you have a known format.
*/

/*
scanner.Scan():
- It's particularly useful for reading and processing lines of text
where you don't necessarily know the format in advance.
- It reads one line at a time as a string,
making it convenient for reading and processing text-based data.
- It's more flexible when dealing with unformatted or multi-token input.
*/
