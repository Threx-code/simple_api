package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Threx-code/simple_api/config"
)

func CreateDB(w http.ResponseWriter, r *http.Request) {
	config.CreateDatabase()

	json.NewEncoder(w).Encode("Database created")
}
