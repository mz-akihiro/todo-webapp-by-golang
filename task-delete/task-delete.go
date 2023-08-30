package taskdelete

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-webapp-by-golang/db"
	"todo-webapp-by-golang/model"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete test")

	dbCnt := db.Newdb()

	defer db.CloseDB(dbCnt)

	if r.Method == "DELETE" {
		fmt.Println("DELETE")
	} else {
		fmt.Println("taskdelete - Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not Allowed")
		return
	}

	// task削除リクエスト用構造体を用
	var deleteTask model.Delete

	// 認証用ミドルウェアがリクエストに内包した[user id]を受け取る
	userId := r.Context().Value(model.ContextKey{UserId: "user_id"})
	deleteTask.UserId = int(userId.(float64))

	// リクエストのbodyに含まれてる削除したい[task id]を抜き取る
	if err := json.NewDecoder(r.Body).Decode(&deleteTask); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}
	fmt.Println(deleteTask)

	// [user id]と[task id]を元にデータベースからの削除を行う
	if _, err := dbCnt.Exec("DELETE FROM task_data WHERE id = ? AND userId = ?", deleteTask.DeleteId, deleteTask.UserId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Database error")
		return
	}
}
