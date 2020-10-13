package middleware

import(
	"net/http"
)

var RequestValidator = func(next http.Handler, foo string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("middleware access.......",foo)
	  next.ServeHTTP(w, r)
	})
  }