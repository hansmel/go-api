package fileio

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"hanmel.com/webservice/models"
)

// UserFilename is the name of the file where the user data is stored
// var UserFilename = "/mnt/web-api/users.json"
var UserFilename = "/users.json"

// var UserFilename = "/Users/grey/dev/mnt/web-api/users.json"

// WriteUsers writes several users to file in json format
func WriteUsers(users []*models.User) {
	f, err := os.Create(UserFilename)
	if err != nil {
		fmt.Println(err)
		return

	}
	jsonusers, _ := json.MarshalIndent(users, "", "   ")
	_, err2 := f.Write(jsonusers)
	if err2 != nil {
		fmt.Println(err2)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// printUsers(users)
}

func printUsers(users []*models.User) {
	prettyJSON, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
