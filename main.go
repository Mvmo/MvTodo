package main

import (
	"os"
	"fmt"
)

var (
	databaseFile = "data.json"
	boardMap = make(map[string]Board)
)

func init() {
	if _, err := os.Stat(databaseFile); err != nil {
		fmt.Println(databaseFile + " doesn't exist - Trying to create ->", err)
		if file, err := os.Create(databaseFile); err != nil {
			fmt.Println("Error while creating " + databaseFile + " file -> ", err)
		} else {
			fmt.Println("File successfully created. Trying to write content...")
			if _, err := file.WriteString("{}"); err != nil {
				fmt.Println("Error while trying to write into file -> ", err)
			} else {
				fmt.Println("File successfully initialized with default json string '{}'")
			}
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
}
