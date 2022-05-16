package Conn

import(
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Con() *sql.DB {
	db, err := sql.Open("mysql", "root:toor@/stuff")
	if err != nil {
		log.Print(err.Error())
	}

	sql := "SHOW COLUMNS FROM things"

	rows, err := db.Query(sql)
	if err != nil {
		log.Print(err.Error())
	}

	for rows.Next() {
		var col string
		rows.Scan(&col)
		fmt.Println(col)
	}

	return db
}