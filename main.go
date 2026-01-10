package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

func main() {
	// DB 연결
	con := DbConnect()
	defer con.Close()

	// Handler 선언
	loginViewHandler := func(w http.ResponseWriter, r *http.Request) {
		tmp1 := template.Must(template.ParseFiles("resources/views/login.html"))
		tmp1.Execute(w, nil)
	}

	loginHandler := func(w http.ResponseWriter, r *http.Request) {
		userId := r.PostFormValue("userId")
		userPwd := r.PostFormValue("userPwd")

		var dbPwd string

		err := con.QueryRow("SELECT user_pwd FROM tb_user WHERE user_id = $1", userId).Scan(&dbPwd)

		log.Print("입력한 아아디: " + userId)
		log.Print("입력한 비밀번호: " + userPwd)

		if err == sql.ErrNoRows {
			htmlStr := "<p style='color: red;'>해당하는 아이디의 계정이 존재하지 않습니다.</p>"
			tmp1, _ := template.New("t").Parse(htmlStr)
			tmp1.Execute(w, nil)
			return
		}

		// 쿼리문 오류 검증 로직
		checkError(err)

		if userPwd != dbPwd {
			htmlStr := "<p style='color: red;'>비밀번호가 올바르지 않습니다.</p>"
			tmp1, _ := template.New("t").Parse(htmlStr)
			tmp1.Execute(w, nil)
			return
		}

		htmlStr := "<p style='color: blue;'>로그인 성공.</p>"
		tmp1, _ := template.New("t").Parse(htmlStr)
		tmp1.Execute(w, nil)
	}

	regViewHandler := func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("resources/views/register.html"))
		tmp.Execute(w, nil)
	}

	// 매핑
	http.HandleFunc("/", loginViewHandler)
	http.HandleFunc("/doLogin", loginHandler)
	http.HandleFunc("/register", regViewHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
