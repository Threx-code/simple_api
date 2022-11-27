package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Threx-code/simple_api/config"
	"github.com/Threx-code/simple_api/utils"
	_ "github.com/go-sql-driver/mysql"
)

type CreateUser struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (req *CreateUser) CreateUser(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	config.Connect()
	db := config.GetDB()

	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelfunc()

	hasehPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		fmt.Printf("password hashing error %s", err.Error())
		return
	}

	query := "INSERT INTO users (firstname, lastname, email, password) VALUES (:firstname, :lastname, :email, :password)"
	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		fmt.Printf("error preparing query %s", err.Error())
		return
	}

	stmt.ExecContext(ctx, req.FirstName, req.LastName, req.Email, hasehPassword)

	createUser := &CreateUser{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	json.NewEncoder(w).Encode(createUser)

}
