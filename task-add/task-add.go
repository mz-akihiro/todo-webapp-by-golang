package taskadd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo-webapp-by-golang/db"
	"todo-webapp-by-golang/model"
)

type taskIdJson struct {
	Id int64 `json:"taskId"`
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	dbCnt := db.Newdb()

	defer db.CloseDB(dbCnt)

	//リクエスト用の構造体をmodelから引っ張ってきて用意
	var task model.Task
	// 認証ミドルウェアからリクエストに内包して送られてきたuseridを抜き出す
	userId := r.Context().Value(model.ContextKey{UserId: "user_id"})
	task.UserId = int(userId.(float64))

	// 用意した構造体にリクエストのjsonをデコード
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}
	fmt.Println(task)

	// taskをデータベースに登録
	if result, err := dbCnt.Exec("INSERT INTO task_data (userId, memo) VALUES (?, ?)", task.UserId, task.Memo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("mysql err", err)
		fmt.Fprintln(w, "Database error")
		return
	} else {
		taskId, _ := result.LastInsertId()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("mysql err", err)
			fmt.Fprintln(w, "Database error")
			return
		}

		// taskIdをレスポンスとして返す
		jsonData := taskIdJson{taskId}
		fmt.Println(jsonData)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonData); err != nil {
			log.Println(err)
		}

	}

}
