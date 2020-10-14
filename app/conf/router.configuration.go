package conf

import(
	"github.com/gorilla/mux"
	"fmt"
	"reflect"
	"go-connect/app/controllers"
	"go-connect/app/middleware"
	"go-connect/app/dto"
	"net/http"
	"github.com/rs/cors"
)




var Router = func() *mux.Router{
	router := mux.NewRouter()
	var c = cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedMethods :[]string{"POST", "PUT","GET","DELETE","OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Accept-Encoding","Authorization", "Content-Type", "X-CSRF-Token"},
    	MaxAge: 300,
		Debug: true,
	})


	RegisterHandler := http.HandlerFunc(controllers.Register)
	router.Handle("/api/register",c.Handler(middleware.RequestValidator(RegisterHandler,reflect.TypeOf(dto.UserRequest{})))).Methods("POST")
	fmt.Println("var1 = ", reflect.TypeOf(router)) 
    return router
}