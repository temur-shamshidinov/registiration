package v1

import (
	"log"
	"registiration/models"
	"registiration/pkg/utils"
	"registiration/storage/redis"

	"github.com/gin-gonic/gin"
)

func (h *handler) CheckUser(ctx *gin.Context) {

	var reqBody models.SignUp

	ctx.Bind(&reqBody)

	notexists, err := h.storage.UserRepo().CheckUserNotExists(reqBody.Username, reqBody.Gmail)

	if err != nil {
		log.Println("error", err)
		ctx.JSON(500, err)
		return
	}

	if !notexists {
		ctx.JSON(200, "username or gmail exists")
		return
	}

	ctx.JSON(200, notexists)
}

func (h *handler) SignUp(ctx *gin.Context) {

	var reqBody models.UserAccount

	ctx.Bind(&reqBody)

	otp := utils.GenerateOtp(6)

	err := redis.Set(otp, "")
	if err != nil {
		return
	}

	utils.SendMail([]string{reqBody.Gmail}, otp)

	ctx.JSON(201, "we are sent otp code! to "+reqBody.Gmail)

}

func (h *handler) CheckOtp(ctx *gin.Context) {

	var reqBody models.CheckOtp
	var account models.UserAccount

	ctx.Bind(&reqBody)

	_, err := redis.Get(reqBody.Gmail)

	if err != nil {
		return
	}

	err = h.storage.UserRepo().CreateUser(account)
	if err != nil {
		log.Println("error in creating user:", err)
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, "you are successfully registred")
}
