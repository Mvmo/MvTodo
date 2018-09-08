package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

var (
	databaseFilename = "data.json"
	databaseFile     *os.File
	boardMap         = make(map[string]Board)
)

func init() {
	if _, err := os.Stat(databaseFilename); err != nil {
		fmt.Println(databaseFilename+" doesn't exist - Trying to create ->", err)
		if file, err := os.Create(databaseFilename); err != nil {
			fmt.Println("Error while creating "+databaseFilename+" file -> ", err)
		} else {
			fmt.Println("File successfully created. Trying to write content...")
			databaseFile = file
		}
	}
}

func main() {

	b, _ := ioutil.ReadFile(databaseFilename)
	json.Unmarshal(b, &boardMap)

	for k, v := range boardMap {
		fmt.Println(k)
		for _, t := range v.Tasks {
			fmt.Println(" -> " + t.Title)
			fmt.Println("  - " + t.Description)
		}
	}
}