package conf

import(
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"reflect"
)

var Router = func() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
	})
	fmt.Println("var1 = ", reflect.TypeOf(router)) 
    return router
}