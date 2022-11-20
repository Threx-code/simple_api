package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Threx-code/simple_api/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	blogModel := &models.Blogs{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), &blogModel)
	blogModel.CreateBlog(w)
}
