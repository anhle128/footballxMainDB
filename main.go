package main

import (
	_ "github.com/lib/pq"
	"github.com/anhle/go-crawler-livescore/database"
	"fmt"
	"database/sql"

)

const (
	dbUser     = "root"
	dbPassword = "123456789"
	dbName     = "football-x-dev"
	dbHost     = "localhost"
	dbPort     = 5432
)

func main() {

	 dbInfo := database.Info{Username:dbUser, Password:dbPassword, DB:dbName, Host:dbHost, Port:dbPort}


	 //database.SetDatabaseInfo(dbInfo)
	 //database.GetData("Select * from ahihii do ngoc")

	 //fmt.Println(database.GetDatabaseInfo())

	 //dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	// fmt.Println(dbinfo)
	 db, err := sql.Open("postgres", dbInfo.GetConnectString())
	 checkErr(err)
	 defer db.Close()

	 rows, err := db.Query("SELECT * FROM clubs")
	 checkErr(err)

	 for rows.Next() {
	 	var id int
	 	var name string
	 	var icon sql.NullString
	 	var deleted bool
	 	err = rows.Scan(&id, &name, &icon, &deleted)
	 	checkErr(err)
	 	// fmt.Println("id | name | icon | deleted ")
	 	strIcon := "null"
	 	if icon.Valid {
	 		strIcon = icon.String
	 	}
	 	fmt.Printf("%3v | %8v | %6v | %6v\n", id, name, strIcon, deleted)
	 }

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ([] string) add(value string ){

}
