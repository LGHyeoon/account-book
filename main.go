package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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
		// 260103 내용
		// 프론트에서 JSON으로 값을 넘길 경우에는 아래와 같이 unmarshal 처리를 해서 데이터를 가져와야한다. (기존의 ioutil.ReadAll은 deprecated 됨)
		// 그 외의 경우로 프론트에서 FormData API를 통해 넘길 경우에는 r.PostFormValue("데이터명") 으로 가져온다.
		body, _ := io.ReadAll(r.Body)
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)

		userId := keyVal["userId"]
		userPwd := keyVal["userPwd"]

		var dbPwd string

		err := con.QueryRow("SELECT user_pwd FROM tb_user WHERE user_id = $1", userId).Scan(&dbPwd)

		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			fmt.Fprint(w, `<p style="color: red;">존재하지 않는 계정입니다.</p>`)
			return
		}

		// 쿼리문 오류 검증 로직
		if err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}

		if userPwd != dbPwd {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			fmt.Fprint(w, `<p style="color: red;">비밀번호가 올바르지 않습니다.</p>`)
			return
		}

		w.WriteHeader(http.StatusNoContent) // 204
		fmt.Fprint(w, `<p style="color: blue;">로그인 성공.</p>`)
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
