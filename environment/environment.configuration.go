package environment

import(
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"go-connect/app/utility"
)
var logger = utility.Logger()

var Environment = func() map[string]string {
	// get the ENV in which Go Application is running. It is  defined while running the go server.
	// command to run the Go Application = "ENV=production go run main.go"
	var env = os.Getenv("ENV")
	var appEnv map[string]string
	var err error

	// loading the properties as per the environment
	if env == "" || env == "development" {
		logger.Info.Println("development environment intiated")
		appEnv, err = godotenv.Read("environment/development.env")
	}else if env == "production"{
		logger.Info.Println("production environment intiated")
		appEnv, err = godotenv.Read("environment/production.env")
	}

	if err != nil {
	  	logger.Error.Println(err)
	}
	
	return appEnv
}
	
	 