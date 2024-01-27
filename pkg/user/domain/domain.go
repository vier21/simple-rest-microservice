package domain

import "time"

type User struct {
	Id string `json:"id,omitempty" bson:"id"`
	Username  string `bson:"userName"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	CreatedAt time.Time `bson:"createdAt"`
	Role     string	`bson:role`
}
