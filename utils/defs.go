package utils

import (
	"database/sql"
	"log"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// DbCon connection
var DbCon *sql.DB

// Logger instance
var Logger *log.Logger
