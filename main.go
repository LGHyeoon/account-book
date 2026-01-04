package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	dbConnect()

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

		log.Print(userId + " / " + userPwd)
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
