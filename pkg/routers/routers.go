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
	authorizedAccountRouter.Use(middleware.AuthMiddleware())
	authorizedAccountRouter.GET("/profile", h.Profile)
	authorizedAccountRouter.DELETE("/delete", h.DeleteUser)

	taskRouter := r.Group("/tasks")
	taskRouter.GET("/", h.AllTasks)
	taskRouter.GET("/:id", h.SingleTask)
	authorizedTaskRouter := r.Group("/tasks")
	authorizedTaskRouter.Use(middleware.AuthMiddleware())
	authorizedTaskRouter.POST("/", h.CreateTask)
	authorizedTaskRouter.PATCH("/:id", h.UpdateTask)
	authorizedTaskRouter.DELETE("/:id", h.DeleteTask)
}
