package main

import(
	"net/http"
	"fmt"
	"go-connect/app/conf"
)

var logger = conf.Logger()

func main(){
	
	// Initializing all the routes
	routes := conf.Router()
	
	logger.Info.Println("Starting the application at http://localhost:8000")
	err := http.ListenAndServe(":"+"8000", routes) 
	if err != nil {
		fmt.Print(err)
	}
}