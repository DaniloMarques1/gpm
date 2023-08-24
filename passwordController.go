package main

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type PasswordController interface {
	SavePassword(w http.ResponseWriter, r *http.Request)
	FindKeys(w http.ResponseWriter, r *http.Request)
	Generate(w http.ResponseWriter, r *http.Request)
	FindPassword(w http.ResponseWriter, r *http.Request)
}

type passwordController struct {
	collection *mongo.Collection
}

func NewPasswordController() (PasswordController, error) {
	client, err := Dabase()
	if err != nil {
		return nil, err
	}
	collection := client.Database("gpm").Collection("password")
	return &passwordController{collection}, nil
}

func (p *passwordController) SavePassword(w http.ResponseWriter, r *http.Request) {
}

func (p *passwordController) FindKeys(w http.ResponseWriter, r *http.Request) {
}

func (p *passwordController) Generate(w http.ResponseWriter, r *http.Request) {
}

func (p *passwordController) FindPassword(w http.ResponseWriter, r *http.Request) {
}
