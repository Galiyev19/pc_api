package models

type Admin struct {
	ID           int    `json:"id"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	HashPassword string `json:"hash_password"`
	Role         string `json:"role"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	IsActive     bool   `json:"is_active"`
}
