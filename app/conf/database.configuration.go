package conf

import (
	"gorm.io/gorm"
	"fmt"
	"gorm.io/driver/postgres"
	"go-connect/app/models"
)

var db *gorm.DB

func init(){
	username := env["db_user"]
	password := env["db_pass"]
	dbName := env["db_name"]
	dbHost := env["db_host"]

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	logger.Info.Println(dsn)
	
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	  }), &gorm.Config{})
	if err != nil {
		logger.Error.Println(err)
	}

	
	db = conn

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	db.Debug().AutoMigrate(&models.RegistrationToken{},&models.User{},&models.Role{},&models.Permission{},&models.Post{},&models.PostAttachment{},&models.Comment{},&models.CommentAttachment{},&models.Profile{},&models.ForgotPasswordToken{})
	
	if err != nil {
		logger.Error.Println(err)
	}
	
}

func GetDB() *gorm.DB {
	return db
}