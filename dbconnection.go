package main

import (
	"database/sql"
	"fmt"

	// 이를 선언하지 않으면 unknown driver 오류가 발생한다.
	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	DATABASE = "accountbook"
	PORT     = "2174"
	USER     = "postgres"
	PASSWORD = "seon9053!"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func DbConnect() *sql.DB {
	// 5개의 설정을 다 해줘야 정상적으로 연결이 가능하다.
	var connectionstring string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)

	db, err := sql.Open("postgres", connectionstring)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("DB 연결 성공")

	return db
}
