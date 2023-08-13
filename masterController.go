package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type MasterController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

type masterController struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMasterControler() (MasterController, error) {
	client, err := Dabase()
	if err != nil {
		return nil, err
	}

	collection := client.Database("gpm").Collection("master")

	return &masterController{client, collection}, nil
}

type SignUpMasterDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m *masterController) SignUp(w http.ResponseWriter, r *http.Request) {
	var masterDto SignUpMasterDto
	if err := json.NewDecoder(r.Body).Decode(&masterDto); err != nil {
		ERROR(w, err)
		return
	}

	if len(masterDto.Email) == 0 || len(masterDto.Password) == 0 {
		ERROR(w, NewApiError("Bad request", http.StatusBadRequest))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(masterDto.Password), bcrypt.DefaultCost)
	if err != nil {
		ERROR(w, err)
		return
	}

	master := Master{
		Id:                     uuid.NewString(),
		Email:                  masterDto.Email,
		HashedPassword:         hashedPassword,
		PasswordExpirationDate: time.Now().Add(time.Hour * 2190),
	}

	if _, err := m.collection.InsertOne(context.Background(), master, options.InsertOne()); err != nil {
		ERROR(w, err)
		return
	}

	JSON(w, nil, http.StatusCreated)
	return
}

type SignInMasterRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInMasterResponseDTO struct {
	AccessToken string `json:"access_token"`
}

func (m *masterController) SignIn(w http.ResponseWriter, r *http.Request) {
	var dto SignInMasterRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		ERROR(w, err)
		return
	}

	var master Master
	err := m.collection.FindOne(context.Background(), bson.M{"email": dto.Email}, options.FindOne()).Decode(&master)
	if err != nil {
		ERROR(w, NewApiError("E-mail or password is incorrect", http.StatusNotFound))
		return
	}

	if err := bcrypt.CompareHashAndPassword(master.HashedPassword, []byte(dto.Password)); err != nil {
		ERROR(w, NewApiError("E-mail or password is incorrect", http.StatusNotFound))
		return
	}

	token, err := Generate(master.Id)
	if err != nil {
		ERROR(w, err)
		return
	}

	JSON(w, SignInMasterResponseDTO{AccessToken: token}, http.StatusOK)
	return
}
