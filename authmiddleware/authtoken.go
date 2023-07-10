package authmiddleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"todo-webapp-by-golang/model"

	"github.com/golang-jwt/jwt"
)

func CheckJwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // 次のハンドラーに渡す
		value, err := r.Cookie("token")
		if err != nil {
			fmt.Println("middleware tokenなし")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else {
			fmt.Println("middleware tokenあり")
		}

		// Cookie内のtokenを検証する
		token, err := jwt.Parse(value.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("error")
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			fmt.Println("token error")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else {
			fmt.Println("jwt is clear", token)
			userID := token.Claims.(jwt.MapClaims)["user_id"].(float64) // 	 ペイロードからuserIdを抜き出す
			fmt.Println("user_id = ", userID)
			ctx := context.WithValue(r.Context(), model.ContextKey{UserId: "user_id"}, userID)
			defer next.ServeHTTP(w, r.WithContext(ctx)) // リクエストにuserIdを内包して次のハンドラーに渡す
		}
	}
}
