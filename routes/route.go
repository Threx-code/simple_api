package routes

import (
	"net/http"

	"github.com/Threx-code/simple_api/controllers"
)

func HttpRoute() {
	http.HandleFunc("/create-db", controllers.CreateDB)
	http.HandleFunc("/index", controllers.Index)
	http.HandleFunc("/create/user", controllers.CreateUser)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}
}
