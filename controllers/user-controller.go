package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Threx-code/simple_api/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	userCreator := &models.CreateUser{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Wrong data %s", err.Error())
	}

	json.Unmarshal([]byte(body), &userCreator)

	userCreator.CreateUser(w)

}
