package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const taskFilePath = ".lomzem.taskapp.tasks"

type Task struct {
	Name string
    Completed bool
}

type TaskList []Task

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func makeTaskFile() {
    _, err := os.Stat(taskFilePath)
    if os.IsNotExist(err) {

        file, err := os.Create(taskFilePath)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()

    }
}

func readTasks(filePath string) TaskList {
    fileContents, err := os.ReadFile(filePath)
    checkError(err)

    if len(fileContents) != 0 {
        var jsonData TaskList
        err = json.Unmarshal(fileContents, &jsonData)
        checkError(err)

        return jsonData
    }
    return TaskList{}
}

func (tasks TaskList) listTasks() {
    fmt.Println("Tasks:\n")
    if len(tasks) == 0 {
        fmt.Println("No tasks to show!")
        return
    }

    for i, task := range tasks {
        switch task.Completed {
        case false: {fmt.Printf("%s) %s\n", strconv.Itoa(i), task.Name)}
        case true: {fmt.Printf("%s) [%s]\n", strconv.Itoa(i), task.Name)}
        }
    }
}

func (tasks *TaskList) addTask(taskName string) {
    newTask := Task{
        taskName,
        false,
    }

    *tasks = append(*tasks, newTask)
    tasks.writeFile()
    fmt.Println("Successfully added ", taskName)
    // }
}

func (tasks *TaskList) removeTask(index int) {
    if index < 0 || index >= len(*tasks) {
        fmt.Println("Invalid index!")
        return
    }

    taskName := (*tasks)[index].Name
    *tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
    tasks.writeFile()
    fmt.Println("Sucessfully removed the task:", taskName)
}

func (tasks *TaskList) markCompleted(index int) {
    if index < 0 || index >= len(*tasks) {
        fmt.Println("Invalid index!")
        return
    }

    taskName := (*tasks)[index].Name
    (*tasks)[index].Completed = true
    tasks.writeFile()
    fmt.Printf("Sucessfully marked the task \"%s\" as completed!", taskName)
}

func (tasks *TaskList) markUncompleted(index int) {
    if index < 0 || index >= len(*tasks) {
        fmt.Println("Invalid index!")
        return
    }

    taskName := (*tasks)[index].Name
    (*tasks)[index].Completed = false
    tasks.writeFile()
    fmt.Printf("Sucessfully marked the task \"%s\" as uncompleted", taskName)
}

func (tasks *TaskList) writeFile() {
    taskBytes, err := json.Marshal(tasks)
    checkError(err)
    os.WriteFile(taskFilePath, taskBytes, 0644)
}

func listCommands() {
    fmt.Println("Task App\n")
    fmt.Println("Available commands:")
    fmt.Println("taskapp list")
    fmt.Println("taskapp add <taskName>")
    fmt.Println("taskapp remove <index>")
    fmt.Println("taskapp complete <index>")
    fmt.Println("taskapp uncomplete <index>")
}

func main() {
    makeTaskFile()
    if len(os.Args) == 1 {
        listCommands()
        return
    }

    tasks := readTasks(taskFilePath)

    switch os.Args[1] {

    case "list":
        tasks.listTasks()

    case "add":
        switch len(os.Args) {

        case 2:
            fmt.Println("Task name required!")
            return

        case 3:
            tasks.addTask(os.Args[2])

        default:
            fmt.Println("Invalid number of arguments!")
        }

    case "remove":
        switch len(os.Args) {
        case 2:
            fmt.Println("Index of task required!")
            return

        case 3:
            intInput, err := strconv.Atoi(os.Args[2])
            if err != nil {
                fmt.Println("Error: did not provide an integer index for task!")
            }
            tasks.removeTask(intInput)

        default:
            fmt.Println("Invalid number of arguments!")
        }

    case "complete":
        switch len(os.Args) {
        case 2:
            fmt.Println("Index of task required!")
            return

        case 3:
            intInput, err := strconv.Atoi(os.Args[2])
            if err != nil {
                fmt.Println("Error: did not provide an integer index for task!")
            }
            tasks.markCompleted(intInput)

        default:
            fmt.Println("Invalid number of arguments!")
        }

        case "uncomplete":
            switch len(os.Args) {
            case 2:
                fmt.Println("Index of task required!")
                return

            case 3:
                intInput, err := strconv.Atoi(os.Args[2])
                if err != nil {
                    fmt.Println("Error: did not provide an integer index for task!")
                }
                tasks.markUncompleted(intInput)

            default:
                fmt.Println("Invalid number of arguments!")
            }

        default:
            listCommands()
        }
    }
