package conf

import(
	"github.com/gorilla/mux"
	"fmt"
	"reflect"
	"go-connect/app/controllers"
)

var logger = Logger()

var Router = func() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api/register", controllers.Register)
	fmt.Println("var1 = ", reflect.TypeOf(router)) 
    return router
}