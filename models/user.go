package models

// Userは、ユーザー情報を表す構造体です
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
