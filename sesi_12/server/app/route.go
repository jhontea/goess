package app

import (
	"pencairan_user/server/controller"
	"pencairan_user/server/middleware"
)

func route() {
	router.Use(middleware.CORSMiddleware()) //to enable api request between client and server
	router.POST("/username", controller.Username)
}
