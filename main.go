package main

import (
	"fmt"
	"net/http"

	"hanmel.com/webservice/controllers"
)

func main() {
	fmt.Println("Starting program!")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	fmt.Println("Stopping program!")
}
