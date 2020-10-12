package main

import(
	"net/http"
	"fmt"
	"go-connect/app/conf"
)

func main(){
	
	// Initializing all the routes
	routes := conf.Router()
	
	err := http.ListenAndServe(":"+"8000", routes) 
	if err != nil {
		fmt.Print(err)
	}
}