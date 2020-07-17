package fileio

import (
	"encoding/json"
	"fmt"
	"os"

	"hanmel.com/webservice/models"
)

// ReadUserFile reads data from the provided file
func ReadUserFile() []*models.User {
	fmt.Println("Reading file " + UserFilename)

	file, err := os.Open(UserFilename)
	if err != nil {
		// handle the error here
		return nil
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return nil
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return nil
	}

	var users []*models.User
	json.Unmarshal(bs, &users)
	return users
}
