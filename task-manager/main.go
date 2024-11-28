package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type TaskId int
type TaskStatus int
type Task map[TaskId]TaskContent

const (
	DONE TaskStatus = iota
	IN_PROGRESS
	PENDING
)

type TaskContent struct {
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func checkTaskStatus(ts TaskStatus) string {
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

func (t *TaskContent) deleteTask() {
	/*
	   TODO: logic to delete an existing task by ID. Check if the task exist
       first before deleting, and if it doesn't I should tell the user.
	*/
}

func (t *TaskContent) updateTask() {
	/*
	   TODO: logic to update the status of a task and change the 
       'UpdatedAt' time. It should Unmarshall and check the exist tasks 
       in the JSON file first.
       If the task doesn't exist I should inform the user.
	*/
}
/*
    TODO: When I add a new task I should load previous tasks first and append 
    the new one in order from oldest to newest.
*/
func addTask(newTask Task) {
    currentTime := time.Now().Format(time.ANSIC)
	var newId TaskId = TaskId(rand.Intn(100))
	newTask[newId] = TaskContent{
		Description: "Test text for my json.",
		Status:      checkTaskStatus(IN_PROGRESS),
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	byteValue, err := json.Marshal(newTask)
	if err != nil {
		fmt.Printf("Failed to create JSON: %s\n", err)
		return
	}

	err = os.WriteFile("data/test.json", byteValue, 0644)
	if err != nil {
		fmt.Printf("Failed to write JSON file: %s\n", err)
	}
}

func main() {
    task := make(Task)
    fileData, err := os.ReadFile("data/test.json")
    if err != nil {
        fmt.Printf("Failed to read JSON file: %s\n", err)
    }
    taskData := []byte(fileData)
    if err := json.Unmarshal(taskData, &task); err != nil {
        fmt.Printf("Failed to Unmarshall JSON data: %s\n", err)
    }
    fmt.Println(task)
    addTask(task)
}
