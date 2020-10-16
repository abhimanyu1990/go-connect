package controllers

import (
	"go-connect/app/utility"
	"go-connect/app/conf"
)

var logger = utility.Logger()
var db = conf.GetDB()