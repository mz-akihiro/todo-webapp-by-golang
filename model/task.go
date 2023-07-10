package model

type Task struct {
	UserId int
	Memo   string `json:"memo"`
}
