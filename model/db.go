package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type EMPLOYEE struct {
	ID     string
	NUMBER string
}

// DB接続とテーブルを作成する
func main() {
	dsn := GetDBConfig()
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	CheckErr(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckErr(err)

	fmt.Println("Connected!")
}

// DBのdsnを取得する
func GetDBConfig() string {
	// 環境変数を読み込む
	// user := os.Getenv("DB_USER")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	dbName := "todolist"
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", hostname, user, password, port, dbName)
	return dsn
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
