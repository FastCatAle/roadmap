package main

import (
	"encoding/json"
	//"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type TaskStatus int
var Task map[int]TaskContent

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

/*func getTaskId(t Task) {
    for k, v := range Task {
        if Task[k] {
            //
        }
    }
}*/

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
func addTask(newTask map[int]TaskContent, description *string) {
    currentTime := time.Now().Format(time.ANSIC)
	var newId int = rand.Intn(100)
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

func main() {
    var Task map[int]TaskContent
    //task := make(Task)
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

    /*newTask := flag.String("add", "test", "Add a new task")
    //flagDeleteTask := flag.Int("delete", 0, "Delete an existing task by ID")
    //flagUpdateTask := flag.Int("update", 0, "Update an existing task by ID")
    //flag.Parse()

    if *newTask != "" {
        fmt.Printf("Flag content %s\n", *newTask)
        addTask(task, newTask)
    } else {
        println("Type '-help' to see all commands")
        os.Exit(0)
    }*/
}
