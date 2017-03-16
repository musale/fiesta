package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/etowett/fiesta/core"
	"github.com/etowett/fiesta/utils"
	"github.com/joho/godotenv"
)

func main() {
	// use dotenv for variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env Error: ", err)
	}

	// Initiate a Log file
	// logFile, err := os.OpenFile(os.Getenv("LOG_FILE"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)

	// if err != nil {
	//     utils.Logger.Fatal("LogFile Error: ", err)
	// }

	// defer logFile.Close()

	// utils.Logger = log.New(logFile, "", log.Lshortfile|log.Ldate|log.Ltime)

	// Mysql Connection
	utils.DbCon, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8")
	if err != nil {
		log.Fatal("DB Error: ", err)
	}
	defer utils.DbCon.Close()

	// Test the connection to the database
	err = utils.DbCon.Ping()
	if err != nil {
		log.Fatal("db ping error: ", err)
	}

	// core.CalcUsage()

	http.HandleFunc("/range", core.RangePage)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
