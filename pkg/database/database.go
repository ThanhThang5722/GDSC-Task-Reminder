package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var Once sync.Once

func GetDbInstance() *sql.DB {
	Once.Do(func() {
		sqlDriverName := os.Getenv("sqlDriverName")
		dataSource := os.Getenv("DB_DSN")
		var err error
		db, err = sql.Open(sqlDriverName, dataSource)
		if err != nil {
			log.Fatalf("Cannot connect to DB")
		} else {
			fmt.Printf("Connect success to DB")
		}
	})
	return db
}
