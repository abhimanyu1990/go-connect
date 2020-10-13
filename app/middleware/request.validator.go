package middleware

import(
	"net/http"
	"io/ioutil"
	"reflect"
	"go-connect/app/dto"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)
var validate = validator.New()
var RequestValidator = func(next http.Handler, dataType reflect.Type) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("middleware access.......")
		reqBody, _ := ioutil.ReadAll(r.Body)
		isValidated, err := objectConverter(reqBody,dataType)

		if isValidated {
			next.ServeHTTP(w, r)
		}else {
			http.Error(w, err.Error(), 404)
		}
		
	 
	})
  }


  func objectConverter(reqBody []uint8, dataType reflect.Type) (bool, error) {
			var objType = dataType.String()
			var isValidated = true
			var err error
			switch objType {
			case "dto.UserRequest":
				obj := dto.UserRequest{}
				json.Unmarshal(reqBody,&obj)
				err = validate.Struct(obj)
				logger.Info.Println("middleware access.......3",err)
			default:
				fmt.Println("two")

			}

			if err != nil {
				isValidated = false
			}
			return isValidated, err
  }