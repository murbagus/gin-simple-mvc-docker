package controllers

import (
	"github.com/gin-gonic/gin"
)

type PingController struct {
	BaseController
}

// Mendefinisikan routing controller
func (controller *PingController) SetRoute(router *gin.RouterGroup) {
	router.GET("/haha", controller.Index)
	router.GET("/panic", controller.Panic)
}

// @Summary Mencoba PING
// @Description Ini hanya dokumentasi percobaan
// @Accept  json
// @Produce  json
// @Param   some_id      path   int     true  "Some ID"
// @Success 200 {string} string	"ok"
// @Router /ping/haha [get]
func (controller *PingController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong-pong",
	})
}

// @Summary Mencoba panic
// @Description Ini hanya dokumentasi percobaan
// @Accept  json
// @Produce  json
// @Param   some_id      path   int     true  "Some ID"
// @Success 200 {string} string	"ok"
// @Router /ping/panic [get]
func (controller *PingController) Panic(c *gin.Context) {
	panic("Panik nich")
}
