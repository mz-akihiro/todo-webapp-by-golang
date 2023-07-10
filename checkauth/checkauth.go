package checkauth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func CheckHandlerJwt(w http.ResponseWriter, r *http.Request) {
	value, err := r.Cookie("token")
	if err != nil {
		fmt.Println("tokenなし1")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println("redirect try")
		return
	} else {
		fmt.Println("tokenあり")
	}

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
	}
}
