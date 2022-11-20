package routes

import (
	"net/http"

	"github.com/Threx-code/simple_api/controllers"
)

func HttpRoute() {
	http.HandleFunc("/index", controllers.Index)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}
}
