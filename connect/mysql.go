package connect

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func MySqlConnect() *sql.DB{
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dbDriver := "mysql"
	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DB")
	dbPort := os.Getenv("MYSQL_PORT")
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, dbConn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}