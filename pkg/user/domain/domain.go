package domain

import "time"

type User struct {
	Id        string    `json:"id,omitempty" bson:"_id"`
	Username  string    `bson:"userName,omitempty"`
	FirstName string    `bson:"firstName,omitempty"`
	LastName  string    `bson:"lastName,omitempty"`
	Email     string    `bson:"email,omitempty"`
	Password  string    `bson:"password,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty"`
	Role      string    `bson:"role,omitempty"`
}
