package config

import (
	"github.com/teris-io/shortid"

	"log"
	"os"
)

var (
	err error  = nil
	PWD string = ""

	// Log Info
	flags = log.Lshortfile
	// Sturct For Information Log
	infoLog = log.New(os.Stdout, "[INFO] ", flags)
	// Sturct For Error Log
	ErrorLog = log.New(os.Stdout, "[ERR] ", flags)
	// Struct For Default Log
	defaultLog = log.New(os.Stderr, "[SYS] ", flags)
)

func Generator() string {
	id, _ := shortid.New(1, shortid.DefaultABC, 2342)
	result, _ := id.Generate()
	return result
}

func SetLog(logType string, msg string) {
	PWD, err = os.Getwd()
	if err != nil {
		ErrorLog.Println("SetLog() -> Couldn't Get Output Of 'os.Getwd()'")
		os.Exit(1)
	}

	logFile, err := os.OpenFile(PWD+"/log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		ErrorLog.Println(`
		Invalid Directory -> /log
		Invalid File -> /log/log.log
		`)
		os.Exit(1)
	}

	switch logType {
	case "I":
		infoLog.SetOutput(logFile)
		infoLog.Println(msg)

	case "E":
		ErrorLog.SetOutput(logFile)
		ErrorLog.Println(msg)

	case "D":
		defaultLog.SetOutput(logFile)
		defaultLog.Println(msg)

	default:
		ErrorLog.SetOutput(logFile)
		ErrorLog.Printf("SetLog() -> Trying To Add Log Without Valid LogType '%s'", logType)
	}

	logFile.Close()
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
	infoLog.SetFlags(log.Ldate | log.Ltime)
	ErrorLog.SetFlags(log.Ldate | log.Ltime)
	SetLog("I", "config.init() -> Flags Setuped")
}
