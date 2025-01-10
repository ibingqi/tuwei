package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {

	ctx.JSON(http.StatusCreated, nil)
}

func Login(ctx *gin.Context) {

	ctx.JSON(http.StatusCreated, nil)
}
