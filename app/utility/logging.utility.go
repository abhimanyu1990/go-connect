package utility

import (
    "log"
    "os"
)

var (
    Warning *log.Logger
    Info    *log.Logger
    Error  *log.Logger
)

type loggers struct{
	Info   *log.Logger
    Warning *log.Logger
    Error   *log.Logger
}

var Logger = func() loggers {
    file, err := os.OpenFile("go-connect.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
    	 log.Fatal(err)
    }


    Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	
	//intiailization as per order is very important
	return loggers{Info,Warning,Error}
	

}