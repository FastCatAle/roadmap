package main

import (
	"encoding/json"
	"flag"
	"fmt"
	//"math/rand"
	"os"
	"time"
)

type TaskStatus int
var Task = make(map[int]TaskContent)

const (
	DONE TaskStatus = iota
	IN_PROGRESS
	PENDING
)

type TaskContent struct {
    //Id    string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func addTaskStatus(ts TaskStatus) string {
    var status string
    switch ts {
    case DONE:
        status = "done"
    case IN_PROGRESS:
        status = "in-progress"
    case PENDING:
        status = "pending"
    }
    return status
}

func getTaskId(task map[int]TaskContent) int{
    var id int = 1
    for k := range task {
        if k == id {
            id++
        }
    }
    return id
}

func deleteTask(task map[int]TaskContent, id int) {
    delete(task, id)
    byteValue, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %s\n", err)
		return
	}

	err = os.WriteFile("data/test.json", byteValue, 0644)
	if err != nil {
		fmt.Printf("Failed to write JSON file: %s\n", err)
	}
    /*
	   TODO: logic to delete an existing task by ID. Check if the task exist
       first before deleting, and if it doesn't I should tell the user.
	*/
}

func updateTask(task map[int]TaskContent, id int, ts TaskStatus) {
    currentTime := time.Now().Format(time.ANSIC)
    switch ts {
    case DONE:
        task[id] = TaskContent{
		    Status:      addTaskStatus(DONE),
            UpdatedAt:   currentTime,
	    }
    case IN_PROGRESS:
        task[id] = TaskContent{
		    Status:      addTaskStatus(IN_PROGRESS),
            UpdatedAt:   currentTime,
	    }
    case PENDING:
        task[id] = TaskContent{
		    Status:      addTaskStatus(PENDING),
            UpdatedAt:   currentTime,
	    }
    }

	byteValue, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %s\n", err)
		return
	}

	err = os.WriteFile("data/test.json", byteValue, 0644)
	if err != nil {
		fmt.Printf("Failed to write JSON file: %s\n", err)
	}

    /*
	   TODO: logic to update the status of a task and change the 
       'UpdatedAt' time. It should Unmarshall and check the exist tasks 
       in the JSON file first.
       If the task doesn't exist I should inform the user.
	*/
}

func addTask(newTask map[int]TaskContent, description *string) {
    currentTime := time.Now().Format(time.ANSIC)
    newId := getTaskId(newTask)
	newTask[newId] = TaskContent{
		Description: *description,
		Status:      addTaskStatus(IN_PROGRESS),
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	byteValue, err := json.Marshal(newTask)
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %s\n", err)
		return
	}

	err = os.WriteFile("data/test.json", byteValue, 0644)
	if err != nil {
		fmt.Printf("Failed to write JSON file: %s\n", err)
	}
}

func readAndUnmarshall() {
    fmt.Println("Reading 'test.json' file...")
    fileData, err := os.ReadFile("data/test.json")
    if err != nil {
        fmt.Printf("Failed to read JSON file: %s\n", err)
    }
    fmt.Println("Unmarshalling 'test.json' file...")
    taskData := []byte(fileData)
    if err := json.Unmarshal(taskData, &Task); err != nil {
        fmt.Printf("Failed to Unmarshall JSON data: %s\n", err)
    }
}

func main() {
    //var Task map[int]TaskContent
    //task := make(Task)
    /*
    fmt.Println("Reading 'test.json' file...")
    fileData, err := os.ReadFile("data/test.json")
    if err != nil {
        fmt.Printf("Failed to read JSON file: %s\n", err)
    }
    fmt.Println("Unmarshalling 'test.json' file...")
    taskData := []byte(fileData)
    if err := json.Unmarshal(taskData, &Task); err != nil {
        fmt.Printf("Failed to Unmarshall JSON data: %s\n", err)
    }
    */
    readAndUnmarshall()
    newTask := flag.String("add", "test", "Add a new task")
    deletedTask := flag.Int("delete", 0, "Delete an existing task by ID")
    updatedTask := flag.Int("update", 0, "Update an existing task by ID")
    flag.Parse()

    if *newTask != "" {
        addTask(Task, newTask)
    } else if updatedTask != nil {
        updateTask(Task, *updatedTask, 0)
    } else if deletedTask != nil {
        deleteTask(Task, *deletedTask)
    } else {
        os.Exit(0)
    }
    
    /*if *newTask != "" {
        fmt.Printf("Flag content %s\n", *newTask)
        addTask(task, newTask)
    } else {
        println("Type '-help' to see all commands")
        os.Exit(0)
    }*/
}
