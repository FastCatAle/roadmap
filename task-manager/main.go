package main

import (
	"encoding/json"
	"fmt"
	"os"
    "time"
)

const (
    DONE  = "done"
    IN_PROGRESS  = "in-progress"
    PENDING = "pending"
)

func (t* Task) deleteTask() {
    /*
    TODO: logic to delete an existing task
    by ID.
    */
}

func (t* Task) updateTask() {
    /*
    TODO: logic to update the status of a update 
    and change the 'UpdatedAt' time.
    */
}

func (t* Task) addTask() {
    //TODO: logic to add a new task with a new ID
}

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
    currentTime := time.Now().Format(time.ANSIC)
	task := []Task{{
		Id:          1,
		Description: "Test text for my json.",
		Status:      IN_PROGRESS,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}}

	byteValue, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("Failed to create JSON: %s\n", err)
		return
	}

	err = os.WriteFile("data/test.json", byteValue, 0644)
	if err != nil {
		fmt.Printf("Failed to write JSON file: %s\n", err)
	}
}
