package services

import(
	"go-connect/app/conf"
	"go-connect/app/utility"
	"go-connect/environment"
)

var db = conf.GetDB()
var logger = utility.Logger()
var env = environment.Environment()