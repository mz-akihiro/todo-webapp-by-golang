package logout

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		fmt.Println("POST")
	} else {
		fmt.Println("logout - Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not Allowed")
		return
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
