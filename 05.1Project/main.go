package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name         string
	description string
	complete     bool
}

type TaskList struct {
	tasks []Task
}

// add a Task
func (tl *TaskList) addTask(t Task) {
	tl.tasks = append(tl.tasks, t)
}

// mark a task as complete
func (tl *TaskList) markAsComplete(index int) {
	tl.tasks[index].complete = true
}

// edit a task
func (tl *TaskList) editTask(index int, task Task) {
	tl.tasks[index] = task
}

// delete a task
func (tl *TaskList) deleteTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func main() {

	// Instance task list
	list := TaskList{}
	read := bufio.NewReader(os.Stdin)
	

	for {
		var option int
		fmt.Println("Choose an option: \n",
			"1. Add Task \n",
			"2. Mark Task as Complete \n",
			"3. Edit Task \n",
			"4. Delete Task \n",
			"5. Exist \n",)
		
		fmt.Println("Select an option: ")
		fmt.Scanln(&option)

		switch option {
		case 1: 
			var task Task
			fmt.Println("Task name: ")
			task.name, _ = read.ReadString('\n')
			
			fmt.Println("Task description: ")
			task.description, _ = read.ReadString('\n')

			list.addTask(task)
			fmt.Println("Task added successfully !")
		case 2: 
			var index int
			fmt.Println("Task index to mark as complete: ")
			fmt.Scanln(&index)
			list.markAsComplete(index)
			fmt.Println("Task marked as complete !")
		case 3: 
			var index int
			var task Task
			fmt.Println("Task index to edit: ")
			fmt.Scanln(&index)

			fmt.Println("Task name: ")
			task.name, _ = read.ReadString('\n')
			
			fmt.Println("Task description: ")
			task.description, _ = read.ReadString('\n')

			list.editTask(index, task)

		case 4: 
			var index int
			fmt.Println("Task index to delete: ")
			fmt.Scanln(&index)

			list.deleteTask(index)
			fmt.Println("Task deleted successfully !")

		case 5: 
			fmt.Println("Goodbye !")
			return

		default: 
			fmt.Println("Invalid option !")
		}

		fmt.Println("Task list :")
		fmt.Println("==============================================")
		
		for i, task := range list.tasks {
			fmt.Printf("%d. - %s - %s Complete: %t\n", i, task.name, task.description, task.complete)
		}
		fmt.Println(("=============================================="))
	}
}
