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
	"github.com/gorilla/context"
	//"github.com/rs/cors"
)
var validate = validator.New()

var RequestValidator = func(next http.Handler, dataType reflect.Type) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("test...")
		reqBody, _ := ioutil.ReadAll(r.Body)
		logger.Info.Println(reqBody)
		isValidated, err := objectConverter(reqBody,dataType,r)
		if isValidated {
			next.ServeHTTP(w, r)
		}else {
			utility.RespondError(w, err.Error(),400)
		}
	})
  }


  func objectConverter(reqBody []uint8, dataType reflect.Type, r *http.Request) (bool, error) {
			var objType = dataType.String()
			var isValidated = true
			var err error
			switch objType {
			case "dto.RegistrationRequest":
				PasswordValidator()
				obj := dto.RegistrationRequest{}
				json.Unmarshal(reqBody,&obj)
				context.Set(r, "reqBody", obj)
				err = validate.Struct(obj)
			default:
				fmt.Println("two")
			}
			if err != nil {
				isValidated = false
			}
			return isValidated, err
  }


  var PasswordValidator = func() {
	_ = validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})
}