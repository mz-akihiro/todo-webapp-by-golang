package logout

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	value, err := r.Cookie("token")
	if err == nil {
		fmt.Println("tokenあり")
		fmt.Println(value)
	} else {
		fmt.Println("tokenなし")
	}

	cookie := &http.Cookie{}

	fmt.Println(cookie.Value)

	cookie.Name = "token"
	cookie.Value = "" // valueを上書きする
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Delete cookie")

}
