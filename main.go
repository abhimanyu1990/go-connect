package main

import(
	"net/http"
	"go-connect/app/conf"
	"go-connect/environment"
	"go-connect/app/utility"
)

var logger = utility.Logger()

func main(){
	
	// Initializing all the routes
	routes := conf.Router()
	env := environment.Environment()
	
	logger.Info.Println("Starting the application at http://localhost:8000 ", env )

	port := env["app_port"]
	if port == "" {
		port = "8000" //localhost
	}
	err := http.ListenAndServe(":"+port, routes) 
	if err != nil {
		logger.Error.Println(err)
	}
}