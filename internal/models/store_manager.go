package models

type StoreManager struct {
	ID           int    `json:"id"`
	UserName     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Password     string `json:"password"`
	HashPassword string `json:"hash_password"`
	Role         string `json:"manager"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	IsActive     bool   `json:"is_active"`
}
