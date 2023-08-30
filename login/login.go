package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"todo-webapp-by-golang/db"
	"todo-webapp-by-golang/model"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	dbCnt := db.Newdb()

	defer db.CloseDB(dbCnt)

	if r.Method == "POST" {
		fmt.Println("POST")
	} else {
		fmt.Println("login - Method not allowed")
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
	fmt.Println("1 pass")

	//userテーブル内のemailカラムに同じ値がないか確認。あればcountが1以上になる。
	//今回はログインなので、emailが登録されていることを確認。
	var count int
	if err := dbCnt.QueryRow("SELECT COUNT(*) FROM user WHERE email = ?", user.Email).Scan(&count); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Database error")
		return
	}
	if count == 0 {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, "Email not registered")
		return
	}
	fmt.Println("2 pass", count)

	//emailの値と同じ行のpasswordカラムから値を取得
	var password string
	row := dbCnt.QueryRow("SELECT password FROM user WHERE email = ?", user.Email)
	if err := row.Scan(&password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Database error")
		return
	}

	//クライアントから送られたパスワードと登録済みのパスワードが一緒かどうか確認
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "different password")
		return
	}

	//tokenにidを埋め込んでおくために取り出す。
	var idToken int
	row2 := dbCnt.QueryRow("SELECT id FROM user WHERE email = ?", user.Email)
	if err := row2.Scan(&idToken); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Database error")
		return
	}
	fmt.Println("ログインしたユーザのID ", idToken)

	//トークン作成
	// パスワードが一致する場合はJWTトークンを発行
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": idToken,                              // 後々taskを追加・削除する際などに使う用に、ペイロードにuserIdを入れる
		"exp":     time.Now().Add(time.Hour * 6).Unix(), // jwtトークンの有効期限(6時間)
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to create token")
		return
	}

	cookie := &http.Cookie{}
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(12 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// ↓Safariで動かない場合はfalseに変更して下さい
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Cookie sent")
}
