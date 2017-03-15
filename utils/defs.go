package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DbCon *sql.DB

var Logger *log.Logger
