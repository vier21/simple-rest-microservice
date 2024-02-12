package domain

import (
	"time"

	"github.com/google/uuid"
)

type JWTPayload struct {
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Roles      []string      `json:"roles"`
	ExpiredAt time.Duration `json:"exp"`
}

type User struct {
	Id        uuid.UUID `json:"id,omitempty" bson:"_id"`
	Username  string    `json:"username" bson:"userName,omitempty"`
	FirstName string    `json:"firstName" bson:"firstName,omitempty"`
	LastName  string    `json:"lastName" bson:"lastName,omitempty"`
	Email     string    `json:"email" bson:"email,omitempty"`
	Password  string    `json:"password" bson:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	Role      string    `json:"role" bson:"role,omitempty"`
}
