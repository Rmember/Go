package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var QUERY_DATA = "SELECT title,price FROM example;"
var (
	ip string = "10.0.2.124"
	user string = "root"
	passwd string = "123"
	port string = "3306"
	database string = "test"
)
func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
func setupConnect() *sql.DB {
	dataSourceName := user + ":" + passwd + "@tcp(" + ip + ":" + port + ")/" + database + "?charset=utf8"
	db, err := sql.Open("mysql",dataSourceName)
	errorHandler(err)
	return db
}
func getdata(db *sql.DB) {
	rows, _ := db.Query(QUERY_DATA)
	defer  rows.Close()
	for rows.Next(){
		var (
			title string
			price int
		)
		err := rows.Scan(&title, &price)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(title,price)
	}
}
func main() {
	db := setupConnect()
	getdata(db)
}
