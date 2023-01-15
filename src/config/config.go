package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sony/sonyflake"
	
	"strconv"
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"
)

var (
	err      error   = nil
	PWD      string  = ""
	Database *sql.DB = &sql.DB{}

	temp     []string = []string{}
	counter  int      = 0
	username string   = ""
	password string   = ""

	flake *sonyflake.Sonyflake = &sonyflake.Sonyflake{}

	// Log Info
	flags = log.Lshortfile
	// Sturct For Information Log
	infoLog = log.New(os.Stdout, "[?] Information -> ", flags)
	// Sturct For Warning Log
	warnLog = log.New(os.Stdout, "[*] Warning -> ", flags)
	// Sturct For Error Log
	errorLog = log.New(os.Stdout, "[!] Error -> ", flags)
	// Struct For Default Log
	defaultLog = log.New(os.Stderr, "[#] Default Log -> ", flags)
)

func Generator() string {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		SetLog("E", "config.Generator() -> Couldn't Generate ID")
		SetLog("D", err.Error())
	}
	return strconv.Itoa(int(id))
}

func SetLog(logType string, msg string) {
	PWD, err = os.Getwd()
	if err != nil {
		errorLog.Println("SetLog() -> Couldn't Get Output Of 'os.Getwd()'")
		os.Exit(1)
	}

	logFile, err := os.OpenFile(PWD+"/log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		errorLog.Println(`
		Invalid Directory -> /log
		Invalid File -> /log/log.log
		`)
		os.Exit(1)
	}

	switch logType {
	case "I":
		infoLog.SetOutput(logFile)
		infoLog.Println(msg)

	case "W":
		warnLog.SetOutput(logFile)
		warnLog.Println(msg)

	case "E":
		errorLog.SetOutput(logFile)
		errorLog.Println(msg)

	case "D":
		defaultLog.SetOutput(logFile)
		defaultLog.Println(msg)

	default:
		errorLog.SetOutput(logFile)
		errorLog.Printf("SetLog() -> Trying To Add Log Without Valid LogType '%s'", logType)
	}

	logFile.Close()
}

func ConnectToSqlDatabase() {

	secret, err := os.Open("./json/Secret.json")
	if err != nil {
		SetLog("E", "config.ConnectToSqlDatabase() -> Can't Open ./json/Secret.json")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}
	defer secret.Close()

	scanner := bufio.NewScanner(secret)
	for scanner.Scan() {
		if scanner.Text() == " " {
			continue
		}

		if (scanner.Text() != "{") && (scanner.Text() != "}") {
			temp = strings.Split(scanner.Text(), "\":\"")

			for i := 0; i < len(temp[1])-1; i++ {
				if counter == 1 {
					password += string(temp[1][i])
				} else {
					if string(temp[1][i]) == "\"" {
						counter += 1
						continue
					}

					username += string(temp[1][i])
				}
			}

		}
	}

	err = scanner.Err()
	if err != nil {
		SetLog("E", "config.ConnectToSqlDatabase() -> Unknow Error From bufio.Scanner()")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}

	Database, err = sql.Open("mysql", username+":"+password+"@tcp(0.0.0.0:3306)/URLShortner")
	if err != nil {
		SetLog("E", "config.ConnectToSqlDatabase() -> Can't Open URLShortner Database")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}
}

func DatabaseStatus() {
	err = Database.Ping()
	if err != nil {
		SetLog("E", "config.DatabaseStatus() -> MySQL Dosn't Response")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}
	SetLog("I", "config.DatabaseStatus() -> MySQL Is Ready To Use")
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
	infoLog.SetFlags(log.Ldate | log.Ltime)
	warnLog.SetFlags(log.Ldate | log.Ltime)
	errorLog.SetFlags(log.Ldate | log.Ltime)
	SetLog("I", "config.init() -> Flags Setuped")
}
