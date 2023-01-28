package sql

import (
	"fmt"
	"os"

	"github.com/JesusKian/URL-Shortner/src/config"
	"github.com/JesusKian/URL-Shortner/src/structure"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"database/sql"
)

var (
	rows      *sql.Rows         = &sql.Rows{}
	Database  *sql.DB           = &sql.DB{}
	MySqlData *structure.Secret = &structure.Secret{}
	err       error             = nil
)

func ReadENV() {
	godotenv.Load("config.env")
	MySqlData = &structure.Secret{
		Username: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
}

func ConnectToSqlDatabase() {
	Database, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(0.0.0.0:3306)/URLShortner", MySqlData.Username, MySqlData.Password))
	if err != nil {
		config.SetLog("E", "sql.ConnectToSqlDatabase() -> Can't Open URLShortner Database")
		config.SetLog("D", err.Error())
		config.ErrorLog.Fatal(err)
	}
}

func DatabaseStatus() {
	err = Database.Ping()
	if err != nil {
		config.SetLog("E", "sql.DatabaseStatus() -> MySQL Dosn't Response")
		config.SetLog("D", err.Error())
		config.ErrorLog.Fatal(err)
	}
	config.SetLog("I", "sql.DatabaseStatus() -> MySQL Is Ready To Use")
}

func init() {
	ReadENV()
	ConnectToSqlDatabase()
	DatabaseStatus()
}
