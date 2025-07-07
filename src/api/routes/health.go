package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omidganji1380/golang-clean-web-api/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("/", handler.Health)
}
