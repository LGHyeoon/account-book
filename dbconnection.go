package main

import (
	"database/sql"
	"fmt"

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

func dbConnect() {
	var connectionstring string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)

	db, err := sql.Open("postgres", connectionstring)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("DB 연결 성공")

	_, err = db.Exec("DROP TABLE IF EXISTS TEST_DB")
	checkError(err)
	fmt.Println("테스트 DB 삭제 완료")

	_, err = db.Exec("CREATE TABLE TEST_DB (id serial PRIMARY KEY, name VARCHAR(50));")
	checkError(err)
	fmt.Println("테스트 DB 생성 완료")

	sql_statement := "INSERT INTO TEST_DB (name) VALUES ($1)"
	_, err = db.Exec(sql_statement, "이기현")
	checkError(err)
	_, err = db.Exec(sql_statement, "이종혁")
	checkError(err)
	fmt.Println("데이터 2개 INSERT 완료")
}
