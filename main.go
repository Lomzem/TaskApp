package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
    TaskPath string `json:"task_path"`
}

type Command struct {
    Name string
    Args string
}

var (
    TaskAdd = Command{"add", "<task name> <due date>"}
    TaskRemove = Command{"remove", "<task name>"}
    TaskChangeDate = Command{"change_date", "<task name> <due date>"}
    TaskList = Command{"list", "list"}
    TaskConfig = Command{"config", "<setting> <new value>"}
    TaskComplete = Command{"complete", "<task name>"}
    TaskUncomplete = Command{"uncomplete", "<task name>"}
)

const DefaultConfigFilePath = ".lomzem.taskapp.config.json"
const DefaultTaskPath = ".lomzem.taskapp.tasks.json"

func makeConfigFile() {
    _, err := os.Stat(DefaultConfigFilePath)
    if os.IsNotExist(err) {
        config := Config{ TaskPath: DefaultTaskPath }
        config_bytes, err := json.Marshal(config)

        if err != nil {
            fmt.Println(err)
        }

        file, err := os.Create(DefaultConfigFilePath)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()

        _, err = file.Write(config_bytes)
        if err != nil {
            fmt.Println(err)
        }
    }
}

func readConfigFile() {
    file, err := os.Open(DefaultConfigFilePath)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    var config Config

    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&config); err != nil {
        fmt.Println(err)
    }

    fmt.Println(config)

}

func changeConfig(config Config) {
}

func list_commands() {
    TaskAdd.printUsage()
    TaskRemove.printUsage()
    TaskChangeDate.printUsage()
    TaskList.printUsage()
    TaskConfig.printUsage()
    TaskComplete.printUsage()
    TaskUncomplete.printUsage()
}

func (cmd Command) printUsage() {
    fmt.Printf("program.exe %s %s", cmd.Name, cmd.Args)
}

func main() {
    // if len(os.Args) == 1 {
    //     list_commands()
    // }
    //
    // switch os.Args[1] {
    // case "add":
    //     args, err := os.Args[:2]
    // default:
    //     fmt.Println("Invalid command!")
    //     list_commands()
    // }
}
