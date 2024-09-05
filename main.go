package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// map which contains the details of the task entered

var taskData = make(map[string]taskDetails)

type taskDetails struct {
	TaskName  string `json:"task name"`
	Tag       string `default:"todo"`
	CreateAt  string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

var count int
var file *os.File

func fileCheck() {
	file, _ = os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0666)
}

func readData() {
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("No data available")
		return
	}

	_ = json.Unmarshal(byteValue, &taskData)
	for id := range taskData {

		count, _ = strconv.Atoi(id)

	}

}


func writeData() {
	serializedData, err := json.MarshalIndent(taskData, "", " ")
	if err != nil {
		fmt.Println("Something Wrong while inserting data")
	}
	// // listTasks()
	file, _ = os.OpenFile("data.json", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	file.Write(serializedData)
}
func addTask(taskName string) {

	fileCheck()
	readData()
	newTask := taskDetails{
		TaskName:  taskName,
		Tag:       "to-do",
		CreateAt:  time.Now().Format("2017.09.07 17:06:06"),
		UpdatedAt: time.Now().Format("2017.09.07 17:06:06"),
	}
	count += 1
	str := fmt.Sprintf("%d", count)
	taskData[str] = newTask
	writeData()
	fmt.Println("Task Added Succesfully")
}

func updateTask(id string, updateTaskName string) {
	fileCheck()
	defer file.Close()
	readData()

	var currentTask taskDetails = taskData[id]
	currentTask.TaskName = updateTaskName
	currentTask.UpdatedAt = time.Now().Format("2017.09.07 17:06:06")

	taskData[id] = currentTask
	writeData()

	fmt.Println("Task Updated Succesfully")

}

func deleteTask(id string) {
	fileCheck()
	defer file.Close()
	readData()
	delete(taskData, id)
	writeData()
	fmt.Println("Deleted Succesfully")
}

func printTasks(status string) {
	fmt.Println("id -> TaskName | Status")
	if status != "" {
		for key, value := range taskData {
			if value.Tag == status {
				fmt.Println(key, "->", value.TaskName, "|", value.Tag)
			}

		}
	} else {
		for key, value := range taskData {
			fmt.Println(key, "->", value.TaskName, "|", value.Tag)
		}
	}

}

func listTasks() {

	fileCheck()
	defer file.Close()
	readData()
	printTasks("")

}

func listTasksWithSpecifiStatus(status string) {
	fileCheck()
	defer file.Close()
	readData()
	printTasks(status)
}

func markProgress(status string, id string) {
	fileCheck()
	defer file.Close()
	readData()

	var currentTask taskDetails = taskData[id]

	fmt.Println("Current Task", currentTask)
	currentTask.Tag = status
	taskData[id] = currentTask
	writeData()
	fmt.Println("Status Updated Succesfully")
}

func main() {
	// root command
	rootCmd := &cobra.Command{
		Use: "task-cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to CLI Application")
		},
	}
	// sub command to add a task to the map data
	addCmd := &cobra.Command{
		Use: "add [task name]",
		Run: func(cmd *cobra.Command, args []string) {
			addTask(args[0])
		},
	}

	// sub command to update a task in the map data
	updateCmd := &cobra.Command{
		Use:  "update [task id] [task name]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			updateTask(args[0], args[1])
		},
	}

	deleteCmd := &cobra.Command{
		Use: "delete [task id]",
		Run: func(cmd *cobra.Command, args []string) {
			deleteTask(args[0])
		},
	}

	listCmd := &cobra.Command{
		Use: "list [status]",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && args[0] != "" {

				listTasksWithSpecifiStatus(args[0])
			} else {
				listTasks()

			}

		},
	}

	// specificTaskCmd := &cobra.Command{
	// 	Use: "[status]",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		listTasksWithSpecifiStatus(args[0])
	// 	},
	// }

	// markingCmd := &cobra.Command{
	// 	Use: "mark-[condition] [id of the task]",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		if args[0] == "in-progress"{
	// 			markProgress("in-progress",args[1])
	// 		}else if args[0] == "done" {
	// 			markProgress("done",args[1])
	// 		}else{
	// 			fmt.Println("Wrong Command")
	// 		}
	// 	},
	// }

	// markingCmd := &cobra.Command{
	// 	Use:   "mark-[in-progress|done] [task ID]",
	// 	Short: "Marks the task as in-progress or done",
	// 	Args:  cobra.ExactArgs(1), // One argument that includes the whole command with a dash
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		taskCmd := args[0]

	// 		if len(args) < 2 {
	// 			fmt.Println("Error: You must provide a task ID")
	// 			return
	// 		}

	// 		taskID := args[1]

	// 		if strings.HasPrefix(taskCmd, "mark-in-progress") {
	// 			markProgress("in-progress", taskID)
	// 		} else if strings.HasPrefix(taskCmd, "mark-done") {
	// 			markProgress("done", taskID)
	// 		} else {
	// 			fmt.Println("Wrong Command: Use 'mark-in-progress [task ID]' or 'mark-done [task ID]'")
	// 		}
	// 	},
	// }

	var markInProgressCmd = &cobra.Command{
		Use:   "mark-in-progress [task ID]",
		Short: "Mark a task as in-progress",
		Args:  cobra.ExactArgs(1), // Ensure exactly one argument for the task ID
		Run: func(cmd *cobra.Command, args []string) {
			taskID := args[0]
			markProgress("in-progress", taskID)
		},
	}

	var markDoneCmd = &cobra.Command{
		Use:   "mark-done [task ID]",
		Short: "Mark a task as done",
		Args:  cobra.ExactArgs(1), // Ensure exactly one argument for the task ID
		Run: func(cmd *cobra.Command, args []string) {
			taskID := args[0]
			markProgress("done", taskID)
		},
	}
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markDoneCmd)
	// listCmd.AddCommand(specificTaskCmd)

	rootCmd.Execute()

}
