package models

import "github.com/golang-jwt/jwt/v4"

type InputRequest struct {
	Email    string `json:"username"	binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token            string `json:"token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresAt        int64  `json:"expires_at"`
	RefreshExpiresAt int64  `json:"refresh_expires_at"`
	Role             string `json:"role"`
}

type TokenClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
