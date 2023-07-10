package main

import (
	"log"
	"net/http"
	"todo-webapp-by-golang/authmiddleware"
	"todo-webapp-by-golang/login"
	"todo-webapp-by-golang/logout"
	"todo-webapp-by-golang/signup"
	taskadd "todo-webapp-by-golang/task-add"
	taskdelete "todo-webapp-by-golang/task-delete"
	taskreload "todo-webapp-by-golang/task-reload"
)

func main() {
	main := http.NewServeMux()
	main.HandleFunc("/signup-api", signup.SignupHandler)
	main.HandleFunc("/login-api", login.LoginHandler)
	main.HandleFunc("/logout-api", logout.LogoutHandler)

	main.Handle("/addtask-api", authmiddleware.CheckJwtMiddleware(taskadd.AddTaskHandler))
	main.Handle("/deletetask-api", authmiddleware.CheckJwtMiddleware(taskdelete.DeleteTaskHandler))
	main.Handle("/reloadtask-api", authmiddleware.CheckJwtMiddleware(taskreload.ReloadTaskHandler))

	main.Handle("/", http.FileServer(http.Dir("public/"))) // 静的ファイル配信。実際にはアパッチなどwebサーバに任せるのが理想。

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", main))
}
