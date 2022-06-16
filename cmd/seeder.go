package main

import (
	"os"
	"encoding/json"
	"log"
	"fmt"

	"go-gin-crud-mysql/model"
	"go-gin-crud-mysql/util"
)

func main () {
	DoSeed()
}

func DoSeed () error {
	// DROP TABLE posts
	u := model.Post{}
	u.DropTable()
	// Create TABLE posts
	u.CreateTable()

	// Get current directory
	var stcData []*model.Post
	p, _ := os.Getwd()
	fmt.Println(p)
	
	// Load local json file
	f, err := os.ReadFile("cmd/n46_members.json")
	if err != nil {
		log.Fatal(err)
	}

	// Read stream from json data
	if err := json.Unmarshal([]byte(f), &stcData); err != nil {
		fmt.Println(err)
		return err
	}

	// Put each row and store into the SQL table
	for _, postRow := range stcData {
		post := model.Post {
			Ulid: util.GenerateULID(),
			Name: postRow.Name,
			Age: postRow.Age,
			Bloodtype: postRow.Bloodtype,
			Origin: postRow.Origin,
		}
		
		// Insert into SQL
		post.Create()
		fmt.Printf("Written data: %s %v\n", postRow.Name, postRow.Age)
	}

	return nil
}
