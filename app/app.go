package app

import (
	"net/http"

	"github.com/nvdhunter/golang-mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
