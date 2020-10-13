package middleware

import(
	"net/http"
	"io/ioutil"
	"reflect"
	"go-connect/app/dto"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-connect/app/utility"
)
var validate = validator.New()
var RequestValidator = func(next http.Handler, dataType reflect.Type) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		isValidated, err := objectConverter(reqBody,dataType)

		if isValidated {
			next.ServeHTTP(w, r)
		}else {
			utility.RespondError(w, err.Error(),400)
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
			default:
				fmt.Println("two")
			}
			if err != nil {
				isValidated = false
			}
			return isValidated, err
  }