package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinisaos/gin-vue-starter/pkg/models"
	"github.com/sinisaos/gin-vue-starter/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Register(c *gin.Context) {

	var userData models.UserRegister

	if err := c.ShouldBind(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userData.Password != userData.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password didn't match."})
		return
	}

	user := models.User{
		UserName: userData.UserName,
		Email:    userData.Email,
		Password: userData.Password,
	}
	user.HashPassword()

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User already exists."})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User created"})
	}
}

func (h *Handler) Login(c *gin.Context) {

	var userData models.UserLogin
	var user models.User
	var err error

	if err := c.ShouldBind(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Model(models.User{}).Where("email=?", userData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email or password is not correct"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email or password is not correct"})
		return
	}

	token, err := utils.GenerateToken(int(user.ID))

	if err != nil {
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("access_token", "Bearer "+token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}