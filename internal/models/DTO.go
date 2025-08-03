package models

type InputRequest struct {
	Email    string `json:"username"	binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
