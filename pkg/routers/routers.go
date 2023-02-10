package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sinisaos/gin-vue-starter/pkg/handlers"
	"github.com/sinisaos/gin-vue-starter/pkg/middleware"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handlers.Handler{
		DB: db,
	}
	// Routes
	accountRouter := r.Group("/accounts")
	accountRouter.POST("/register", h.Register)
	accountRouter.POST("/login", h.Login)
	accountRouter.POST("/logout", h.Logout)
	authorizedAccountRouter := r.Group("/accounts")
	authorizedAccountRouter.Use(middleware.JwtAuthMiddleware())
	authorizedAccountRouter.GET("/profile", h.Profile)
	authorizedAccountRouter.DELETE("/delete", h.DeleteUser)
}
