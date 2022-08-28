package models

type UserInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Created string `db:"created_at"`
	Updated string `db:"updated_at"`
}
