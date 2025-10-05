package main

type Task struct {
	name         string
	descriptions string
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

}
