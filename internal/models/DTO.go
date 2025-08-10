package models

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
	UserID         int    `json:"user_id"`
	Role           string `json:"role"`
	StandardClaims `json:"standard_claims"`
}

type StandardClaims struct {
	ExpiresAt int64 `json:"expires_at"`
	IssuedAt  int64 `json:"issued_at"`
}
