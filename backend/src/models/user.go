package models

type User struct {
	Username       string `json:"username" bson:"username"`
	HashedPassword string `json:"hashed_password" bson:"hashed_password"`
	Email          string `json:"email" bson:"email"`
	FirstName      string `json:"first_name" bson:"first_name"`
	LastName       string `json:"last_name" bson:"last_name"`
	IsThirdParty   bool   `json:"is_third_party" bson:"is_third_party"`
}
