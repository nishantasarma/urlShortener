package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func initConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:helloworld@tcp(localhost:3306)/urlmap")

	if err != nil {
		log.Fatal(err)
	}

	return db

}

func Insertdb(shortUrl string, longUrl string) {

	db := initConnection()
	defer db.Close()
	tx, _ := db.Begin()

	stmt, err1 := tx.Prepare("insert into urls values(?,?)")
	if err1 != nil {

		log.Fatal(err1)

	}
	result, err2 := stmt.Exec(&shortUrl, &longUrl)
	if err2 != nil {
		tx.Rollback()
		log.Fatal(err2)

	}
	tx.Commit()
	fmt.Println(result)

}

var u = map[string]string{}

func Getdb() *map[string]string {

	var shortUrl, longUrl string

	db := initConnection()
	defer db.Close()
	stmt, err := db.Prepare("select * from urls")
	if err != nil {

		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&shortUrl, &longUrl)
		if err != nil {
			log.Fatal(err)

		}
	}
	u[shortUrl] = longUrl

	if u == nil {
		fmt.Println("Database is empty")
		os.Exit(1)
	}
	return &u
}
