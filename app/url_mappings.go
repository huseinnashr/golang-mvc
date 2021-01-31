package app

import (
	"github.com/nvdhunter/golang-mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
