package conf

import(
	"github.com/gorilla/mux"
	"fmt"
	"reflect"
	"go-connect/app/controllers"
	"go-connect/app/middleware"
	"go-connect/app/dto"
	"net/http"
)

var Router = func() *mux.Router{
	router := mux.NewRouter()
	
	RegisterHandler := http.HandlerFunc(controllers.Register)
	router.Handle("/api/register", middleware.RequestValidator(RegisterHandler,reflect.TypeOf(dto.UserRequest{}))).Methods("POST")
	fmt.Println("var1 = ", reflect.TypeOf(router)) 
    return router
}