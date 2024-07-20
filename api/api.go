package api

import (
	v1 "registiration/api/handlers/v1"
	"registiration/storage"

	"github.com/gin-gonic/gin"
)

func Api(storage storage.StorageI) {

	router := gin.Default()

	h := v1.NewHandler(storage)

	router.GET("/ping", h.Ping)

	group := router.Group("/v1")
	{
		group.POST("/check-user", h.CheckUser)
		group.POST("/sign-up", h.SignUp)
		group.POST("/check-otp", h.CheckOtp)

	}

	router.Run(":8080")
}
