package web

import "time"

type UserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct { 
	Status       string `json:"status"`
	Code         int `json:"code"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt time.Duration `json:"expired_at"`
	Role         string
}
