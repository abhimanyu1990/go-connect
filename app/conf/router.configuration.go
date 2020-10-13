package conf

import(
	"github.com/gorilla/mux"
	"fmt"
	"reflect"
	"go-connect/app/controllers"
	"go-connect/app/middleware"
	"net/http"
)

var Router = func() *mux.Router{
	router := mux.NewRouter()
	
	RegisterHandler := http.HandlerFunc(controllers.Register)
	var test = "testing"
	router.Handle("/api/register", middleware.RequestValidator(RegisterHandler, test))
	fmt.Println("var1 = ", reflect.TypeOf(router)) 
    return router
}