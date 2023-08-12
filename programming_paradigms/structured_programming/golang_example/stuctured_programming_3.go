package golang_example

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Task struct {
	Title     string
	Completed bool
}

type TaskManager struct {
	Tasks        []Task
	DataFilePath string
}

func NewTaskManager(dataFilePath string) *TaskManager {
	tasks := loadTasks(dataFilePath)
	return &TaskManager{
		Tasks:        tasks,
		DataFilePath: dataFilePath,
	}
}

func (tm *TaskManager) AddTask(title string) {
	task := Task{Title: title, Completed: false}
	tm.Tasks = append(tm.Tasks, task)
	saveTasks(tm.DataFilePath, tm.Tasks)
	fmt.Println("Task added:", title)
}

func (tm *TaskManager) MarkTaskAsCompleted(index int) {
	if index >= 0 && index < len(tm.Tasks) {
		tm.Tasks[index].Completed = true
		saveTasks(tm.DataFilePath, tm.Tasks)
		fmt.Println("Task marked as completed:", tm.Tasks[index].Title)
	} else {
		fmt.Println("Invalid index.")
	}
}

func (tm *TaskManager) ListTasks() {
	fmt.Println("To-Do List:")
	for i, task := range tm.Tasks {
		status := " "
		if task.Completed {
			status = "X"
		}
		fmt.Printf("%d. [%s] %s\n", i, status, task.Title)
	}
}

func loadTasks(dataFilePath string) []Task {
	var tasks []Task
	fileData, err := ioutil.ReadFile(dataFilePath)
	if err == nil {
		decoder := gob.NewDecoder(strings.NewReader(string(fileData)))
		err := decoder.Decode(&tasks)
		if err != nil {
			log.Println("Error decoding tasks:", err)
		}
	}
	return tasks
}

func saveTasks(dataFilePath string, tasks []Task) {
	file, err := os.Create(dataFilePath)
	if err != nil {
		log.Println("Error creating data file:", err)
		return
	}
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		log.Println("Error encoding tasks:", err)
	}
	file.Close()
}

func main() {
	dataFilePath := "tasks.dat"
	taskManager := NewTaskManager(dataFilePath)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as Completed")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &choice)

		switch choice {
		case 1:
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := scanner.Text()
			taskManager.AddTask(title)
		case 2:
			fmt.Print("Enter task index to mark as completed: ")
			var index int
			fmt.Sscan(scanner.Text(), &index)
			taskManager.MarkTaskAsCompleted(index)
		case 3:
			taskManager.ListTasks()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}
