# Task-CLI
A simple Command-Line Interface (CLI) task manager built in Go, allowing users to create, view, and manage tasks effectively.

## Features
Add tasks with names and tags.
View tasks with specific tags.
Simple and lightweight with a focus on ease of use.
## Requirements
* Go version 1.16 or later
* Cobra package (github.com/spf13/cobra)

## Installation
To install the Task-CLI application, follow these steps:
1. Clone the repository:
```
git clone https://github.com/shanmukha2491/Task-Tracker.git
cd task-cli
go mod init taskmanager
```
2. Install the dependencies

```
go mod tidy
```

3. Build the application

```
go build -o task-cli
```
4. Change the directory of package task-cli

```
echo $PATH
```

### copy the path until `bin` without semicolon and use the following command

```
sudo mv task-cli <path>
```


# Usage
Here are the commands you can use with Task-CLI:

## Initialize the CLI
Run the following command to initialize the application:
```
task-cli
```

1. Add a task
```
task-cli add "task name"
```
2. Delete a task

```
task-cli delete <id of task>
```

3. Update a task
```
task-cli update <id> <new task name>
```
4. Get all tasks

```
task-cli list
```

5. Get task with specific status
```
task-cli list done
task-cli list todo
task-cli list in-progress
```
6. Update the status of the task
```
task-cli mark-in-progress <task id>
task-cli mark-done <task id>
```

Contribution
Feel free to open issues or pull requests for improvements or bug fixes.




