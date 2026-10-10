package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"fake.com/instore/client"
	"fake.com/instore/storepb"
	"google.golang.org/grpc"
)

const (
	bCmd string = "[cmd]  "
	bUsr string = "[usr]> "
)

// if exitSign is present, the program terminates
var exitSigns = []string{"exit", "quit", "bye", "goodbye"}

// commands available to be used
var commands = []string{"use", "add", "addcsv", "get", "getall", "getk", "remove", "removecsv"}

// if a dbsign is present after command `use` the following commands are done on MongoDB otherwise on csvfile
var dbSigns = []string{"db", "database", "data-base", "mongo", "mongod", "mongodb"}

// storeType = "db" means that we are going to use mongodb, otherwise csv
var storeType = "csv"

func newConnection() *grpc.ClientConn {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	return cc
}

func main() {
	// initializeing filePath to for later use
	var _, b, _, _ = runtime.Caller(0)
	// basepath = "/path/to/command" the directory in which we find command.go
	var basepath = filepath.Dir(b)
	// filePath = /path/to/command/csvFiles/ the directory in which we find csv file to be used by client
	var filePath = basepath + "/csvFiles/"
	_ = filePath

	cc := newConnection()
	c := storepb.NewStoreServiceClient(cc)

	defer cc.Close()

	// First screen to be shown to the user
	fmt.Println(bCmd + "Hello! I am excited to have you as a user :D Enjoy!")
	fmt.Println(bCmd + "You could use any of these commands")
	fmt.Println(bCmd + "----------------------------------------------------------------------------------------------")
	fmt.Println(bCmd + "add       <key>  <value>  ------- to add a key-value pair to the store")
	fmt.Println(bCmd + "get[v]    <key>           ------- to get existing values corresponding to a key from the store")
	fmt.Println(bCmd + "getk      <value>         ------- to get existing keys corresponding to a value from the store")
	fmt.Println(bCmd + "remove    <key>  <value>  ------- to remove a key-value pair from the store")
	fmt.Println(bCmd + "exit/quit                 ------- to quit the program")
	fmt.Println(bCmd + "----------------------------------------------------------------------------------------------")

	// infinite loop (command line)
	for {
		// Get input from User
		fmt.Print(bUsr)
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')

		// in case error happend while reading user input
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		// Trim leading and trailing spaces and new-line char
		cmdString = strings.Trim(cmdString, " \n")

		// split input into keywords
		cmdKeys := strings.Split(cmdString, " ")

		// lowercase the command
		cmdKeys[0] = strings.ToLower(cmdKeys[0])

		// similar commands
		switch cmdKeys[0] {
		case "delete":
			cmdKeys[0] = "remove"
		case "getv":
			cmdKeys[0] = "get"
		default:
		}

		// check for exiting tokens
		if contains(cmdKeys[0], exitSigns) {
			fmt.Println(bCmd + "Bye Bye! :)")
			return
		}

		// check for length of the input
		if len(cmdKeys) > 3 {
			fmt.Println(bCmd + "Usage: <command> <key> <value>")
			continue
		}

		// check for empty line
		if strings.Trim(cmdKeys[0], " ") == "" {
			continue
		}

		switch cmdKeys[0] {
		case "use":
			if len(cmdKeys) != 2 {
				fmt.Println(bCmd + "Usage: use [csv]")
			} else if cmdKeys[1] == "csv" {
				client.Use(c, "csv")
			}
		case "insert":
			if len(cmdKeys) != 3 {
				fmt.Println(bCmd + "Usage: insert <key> <value>")
			} else {
				client.Insert(c, cmdKeys[1], cmdKeys[2])
			}
		case "select":
			if cmdKeys[1] == "*" {
				client.Select(c, "*")
			}
		default:
			fmt.Println(bCmd + "'" + string(cmdKeys[0]) + "' is not a command!")
			fmt.Println(bCmd + "Please use one of the commands provided above!")
		}
		

	}
}

// contains is a function that checks whether token is in s or not
func contains(token string, s []string) bool {
	for _, v := range s {
		if strings.Contains(token, v) {
			return true
		}
	}
	return false
}
