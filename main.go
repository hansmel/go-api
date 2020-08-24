package main

import (
	"fmt"
	"net/http"

	"hanmel.com/webservice/models"

	"hanmel.com/webservice/controllers"
	"hanmel.com/webservice/fileio"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)
func main() {


	fmt.Println("Starting program!")

	// Populate users from file
	users := fileio.ReadUserFile()
	models.InitUsers(users)

	// Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	fmt.Println("Stopping program!")
}
