package fileio

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

// FileWriter writes data to file
func FileWriter(user models.User) {
	fmt.Println("Filewriter")
	fmt.Println(user)
}