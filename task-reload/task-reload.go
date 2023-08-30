package taskreload

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo-webapp-by-golang/db"
	"todo-webapp-by-golang/model"
)

func ReloadTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reload test")

	dbCnt := db.Newdb()

	defer db.CloseDB(dbCnt)

	if r.Method == "POST" {
		fmt.Println("POST")
	} else {
		fmt.Println("taskreload - Method not Allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not Allowed")
		return
	}

	// 認証用ミドルウェアがリクエストに内包した[user id]を受け取る
	userId := r.Context().Value(model.ContextKey{UserId: "user_id"})

	// [taskId, memo]を受け取るようのスライスを用意する。
	var taskData []model.ReloadTask

	// データベースから自身のidで作成されたtaskを抜き出す
	if rows, err := dbCnt.Query("SELECT id,memo FROM task_data WHERE userId = ?", int(userId.(float64))); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("task-reload database err", err)
		fmt.Fprintln(w, "Database error")
		return
	} else {
		defer rows.Close()
		for rows.Next() { // 抜き出したtaskをイテレーターで回してtaskDataに格納する
			var td model.ReloadTask
			if err := rows.Scan(&td.TaskId, &td.Memo); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("task-reload database err", err)
				fmt.Fprintln(w, "Database error")
				return
			}
			taskData = append(taskData, td)
		}
		if err := rows.Err(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("task-reload database err", err)
			fmt.Fprintln(w, "Database error")
			return
		}
	}
	fmt.Println("reload task data", taskData)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskData); err != nil { // jsonをセットしてレスポンスとして返す
		log.Println(err)
	}

}
