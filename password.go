package main

type Password struct {
	Id       string `bson:"_id"`
	MasterId string `bson:"master_id"`
	Key      string `bson:"key"`
	Value    string `bson:"value"`
}

type PasswordRepository interface {
	Save(*Password) error
	FindByKey(string) (*Password, error)
	FindByMaster(string) ([]Password, error)
}
