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
	/*srv := echo.New()

	srv.Static("/", "public/")

	apiGroup := srv.Group("/api")
	apiV1Group := apiGroup.Group("/v1")
	apiV1Group.GET("/test", func(context echo.Context) error {
		return context.JSON(200, "IT WORKS")
	})

	srv.Start(":8149")*/

	/*board := Board{Password: "foo", UniqueId: "entendorf", Tasks: []Task{
		{
			Title: "Bug fix #2",
			Description: "Da ist ein Fehler",
		},
		{
			Title: "Bug fix #1",
			Description: "Da ist ein Fehler 2ea",
		},
		{
			Title: "Bug fix #5",
			Description: "Da ist ein Fehler daw",
		},
	}}

	boardMap["entendorf"] = board

	b, _ := json.Marshal(boardMap)
	databaseFile.Write(b)*/

	b, _ := ioutil.ReadFile(databaseFilename)
	json.Unmarshal(b, &boardMap)

	//fmt.Println(boardMap)
	for k, v := range boardMap {
		fmt.Println(k)
		for _, t := range v.Tasks {
			fmt.Println(" -> " + t.Title)
			fmt.Println("  - " + t.Description)
		}
	}

	/*c, _ := ioutil.ReadFile(databaseFilename)
	var j = make(map[string]interface{})
	json.Unmarshal(c, &j)*/
}