package main

import (
	"github.com/gin-gonic/gin"
	"oceanlearn.teach/ginessential/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	return r
}
