package signup

import (
	"encoding/json"
	"fmt"
	"net/http"

	"todo-webapp-by-golang/db"
	"todo-webapp-by-golang/model"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	dbCnt := db.Newdb()

	defer db.CloseDB(dbCnt)

	if r.Method == "POST" {
		fmt.Println("POST")
	} else {
		fmt.Println("signup - Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not Allowed")
		return
	}

	//リクエスト用の構造体をmodelから引っ張ってきて用意
	var user model.User
	//用意した構造体にリクエストのjsonをデコード
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}
	fmt.Println(user)

	//email passwordにちゃんと値が入ってるか確認
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Email and password are required")
		return
	}

	//userテーブル内のemailカラムに同じ値がないか確認。あればcountが1以上になる。
	//今回は初回登録なので、すでに登録されていたらアウト
	var count int
	if err := dbCnt.QueryRow("SELECT COUNT(*) FROM user WHERE email = ?", user.Email).Scan(&count); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Database error")
		return
	}
	if count > 0 {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, "Email already exists")
		return
	}

	//bcryptでハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Hashing error")
		return
	}
	if _, err := dbCnt.Exec("INSERT INTO user (email, password) VALUES (?, ?)", user.Email, hash); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Database error")
		return
	}
}
