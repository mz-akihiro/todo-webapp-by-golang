package main

import (
	"fmt"
	"log"
	"todo-webapp-by-golang/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbCnt := db.Newdb()

	defer db.CloseDB(dbCnt)

	_, err := dbCnt.Exec(`CREATE TABLE IF NOT EXISTS task_data (
						id INT AUTO_INCREMENT PRIMARY KEY,
						userId INT NOT NULL,
						memo VARCHAR(255) NOT NULL,
						created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
						updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
					)`)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("task_data table created successfully")
}
