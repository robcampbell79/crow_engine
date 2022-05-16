package Select

import(
	//"database/sql"
	"fmt"
	"log"
	"crow_engine/Conn"

	_ "github.com/go-sql-driver/mysql"
)

func SelectAll() {
	db := Conn.Con()

	defer db.Close()

	sql := "SELECT item FROM things"

	selectAll, err := db.Query(sql)
	if err != nil {
		log.Print(err.Error())
	}

	for selectAll.Next() {
		var col string
		selectAll.Scan(&col)
		fmt.Println(col)
	}

}