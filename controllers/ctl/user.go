package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hzl.im/gin-platform/models"
	"hzl.im/gin-platform/controllers"
	"strconv"
)

func UserInfo(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	user := &models.User{}
	if err := controllers.DB.First(&user, user_id).Error; err != nil {
		res := models.ResultData{1, "user not exist", ""}
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := models.ResultData{0, "", user}
	ctx.JSON(http.StatusOK, res)
}

func UserAdd(ctx *gin.Context) {
	user := &models.User{}
	if ctx.Bind(user) != nil {
		res := models.ResultData{9999, "params wrong", ""}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := controllers.DB.Create(user).Error; err != nil {
		res := models.ResultData{9998, "add user fail", ""}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := models.ResultData{0, "", ""}
	ctx.JSON(http.StatusOK, res)
}

func UserDel(ctx *gin.Context) {
	user_id := ctx.PostForm("user_id")
	if _, err := strconv.Atoi(user_id); err!=nil {
		res := models.ResultData{9999, "params wrong", ""}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	user := &models.User{}
	if err := controllers.DB.First(&user, user_id).Error; err != nil {
		res := models.ResultData{1, "user not exist", ""}
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	if err := controllers.DB.Delete(&user).Error; err != nil {
		res := models.ResultData{9998, "del user fail", ""}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := models.ResultData{0, "", ""}
	ctx.JSON(http.StatusOK, res)
}