package routes

import (
	"backend/handlers/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	apiGroup := r.Group("/api/v1")
	apiGroup.GET("/auth/google", auth.Google)
	apiGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is healthy"})
	})

	pingGroup := apiGroup.Group("/ping")
	pingRouter(pingGroup)

}
