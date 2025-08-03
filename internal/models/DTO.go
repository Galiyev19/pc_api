package models

type InputRequest struct {
	UserName string `json:"username"	binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name,omitempty"`
}
