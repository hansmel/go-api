package main

import (
	"fmt"
	"net/http"

	"hanmel.com/webservice/models"

	"hanmel.com/webservice/controllers"
	"hanmel.com/webservice/fileio"
)

func main() {
	fmt.Println("Starting program!")

	// Populate users from file
	users := fileio.ReadUserFile()
	models.InitUsers(users)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	fmt.Println("Stopping program!")
}
