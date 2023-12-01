package controllers

import (
	"net/http"

	"github.com/Dparty/common/server"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init(addr ...string) {
	router = gin.Default()
	router.Use(server.CorsMiddleware())
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": "0.0.2",
		})
	})
	router.Run(addr...)
}
