package fileio

import (
	"encoding/json"
	"fmt"
	"os"

	//"hanmel.com/webservice/models"
)

// WriteUsers writes several users to file
//func WriteUsers(users []*models.User, filename string) {
	func WriteUsers(users interface{}, filename string) {

	f, err := os.Create(filename)
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
}

