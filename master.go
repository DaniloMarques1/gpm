package main

import "time"

type Master struct {
	Id                     string    `bson:"_id"`
	Email                  string    `bson:"email"`
	HashedPassword         []byte    `bson:"hashed_password"`
	PasswordExpirationDate time.Time `bson:"password_expiration_date"`
}

type ManagerRepository interface {
	Save(*Master) error
	FindByEmail(string) (*Master, error)
}
