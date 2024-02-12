package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"
    "path/filepath"
)

func main() {
    reader := bufio.NewReader(os.Stdin) // outside the loop, only need one reader

    currentDir, err := os.Getwd() // Get location of current working directory, used for "virtual directory"
    if err != nil {
        os.Exit(1)
    }

    fmt.Println("\nPlease type in a command: 'ls', 'wc', 'mkdir', 'cp', 'mv', 'whoami', 'help', or 'exit'")

    for {
        fmt.Print(currentDir + ": ")
        input, err := reader.ReadString('\n') //get user command

        if err != nil { // check input error
            log.Fatal("Failed to read input:", err)
        }

        input = strings.TrimSpace(input) // remove new lines

        if input == "" { //check for no input (prevent string pars crash)
            fmt.Print("Command not recognized")
            continue
        }

        inputArray := strings.Fields(input) // split inputs by spaces for args
        command := inputArray[0] // command is its own var
        commandArgs := inputArray[1:] // args are its own array

        switch command {
        case "whoami": // prints name
            fmt.Println("Dakotah Proffit, dproffit1")
        case "exit":// exit program
            os.Exit(0)
        case "ls": // list directory, args provided
            var commandArgs []string
            commandArgs = append(commandArgs, currentDir)
            runCommand("ls", commandArgs)
        case "wc": // uses given file
            commandArgs[0] = currentDir + "/" + commandArgs[0] // run arg in virtual directory
            runCommand("wc", commandArgs)
        case "mkdir": // makes given file
            commandArgs[0] = currentDir + "/" + commandArgs[0] // run arg in virtual directory
            runCommand("mkdir", commandArgs)
        case "cp": // copies given file to given location
            commandArgs[0] = currentDir + "/" + commandArgs[0] // run arg in virtual directory
            commandArgs[1] = currentDir + "/" + commandArgs[1] // run arg in virtual directory
            runCommand("cp", commandArgs)
        case "mv": // moves a given file to given location
            commandArgs[0] = currentDir + "/" + commandArgs[0] // run arg in virtual directory
            commandArgs[1] = currentDir + "/" + commandArgs[1] // run arg in virtual directory
            runCommand("mv", commandArgs)
        case "cd": // update workingDirectory
            currentDir = runCD(currentDir, commandArgs)
        case "help":
            fmt.Println("commands: 'ls', 'wc', 'mkdir', 'cp', 'mv', 'whoami', 'help', or 'exit'")
        default:
            fmt.Println("Command not recognized")
        }
    }
}

func runCommand(command string, args []string) {
    cmd := exec.Command(command, args...) //... used for slicing arguments
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("Failed to execute command: ", err)
    }
}

func runCD(workingDirectory string, newLocation []string) string {
    if len(newLocation) == 0 {
        fmt.Println("No directory specified")
        return workingDirectory // Return the current working directory by default
    }
    
    newDirectory := filepath.Join(workingDirectory, newLocation[0])
    if _, err := os.Stat(newDirectory)
    err != nil {
        fmt.Println("CD location not found:", err)
        return workingDirectory // Return the current directory if the new directory does not exist
    }
    return newDirectory // Return the new directory if it exists
}