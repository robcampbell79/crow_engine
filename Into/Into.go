package Into

import(
	//"database/sql"
	"fmt"
	//"log"
	"crow_engine/Conn"

	_ "github.com/go-sql-driver/mysql"
)

func SaveInto(name string) {
	db := Conn.Con()

	defer db.Close()

	sql := "INSERT INTO things(item) VALUES(?)"

	//db.Prepare(sql)

	db.Exec(sql, name)

	fmt.Println("Inserting: ", name)
}