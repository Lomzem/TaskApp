package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type Config struct {
    TaskPath string `json:"task_path"`
}

const DefaultConfigFilePath = ".lomzem.taskapp.config.json"
const DefaultTaskPath = ".lomzem.taskapp.tasks.json"

func make_config_file() {
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

func read_config_file() {
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

func change_config(config Config) {
}

func get_input(reader io.Reader, args ...string) (string, error) {
    fmt.Println(args)
    // if len(args) > 0 {}
    return "", nil
}

func main() {
    config := flag.Bool("config", false, "Change config for task app")
    flag.Parse()

    switch {
    case *config:
        args, err := get_input(os.Stdin, flag.Args()...)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        println(args)
        // fmt.Printf(config)
    }

	// make_config_file()
    // read_config_file()
    // task_app config taskpath string
}
