package fileio

import (
	"os"
	"fmt"

	"github.com/pluralsight/webservice/models"
)

// FileWriter writes user data to file
func FileWriter(user models.User) {
	fmt.Println("Filewriter")

	f, err := os.Create("users.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(user.FirstName + " " + user.LastName + "\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}

