package web

import "github.com/lestrrat-go/jwx/v2/jwt"

type UserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken jwt.Token
}

