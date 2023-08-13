package main

type Password struct {
	Key   string `bson:"password"`
	Value string `bson:"value"`
}

type PasswordRepository interface {
	Save(*Password) error
	FindByKey(string) (*Password, error)
}
