package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sinisaos/gin-vue-starter/pkg/models"
	"github.com/sinisaos/gin-vue-starter/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

// Register User godoc
//
//	@Summary		Register User
//	@Description	Add a new User
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserRegister	true	"User Data"
//	@Success		200		{object}	models.User
//	@Router			/accounts/register [post]
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
		// It should be status 422 (http.StatusUnprocessableEntity) or 409 Conflict
		// but in that case the Axios response is undefined and we can't handle the error
		// if user alredy exists. With curl and swagger docs response is ok.
		c.JSON(http.StatusOK, gin.H{"error": "User already exists."})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User created"})
	}
}

// Login user godoc
//
//	@Summary		Login User
//	@Description	Login User
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserLogin	true	"User Data"
//	@Success		200		{object}	models.UserLogin
//	@Router			/accounts/login [post]
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

// Logout user godoc
//
//	@Summary		User logout
//	@Description	User logout
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/accounts/logout [post]
func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// User delete godoc
//
//	@Summary		User delete
//	@Description	User delete
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Success		204
//	@in				header
//	@name			Authorization
//	@Security		BearerAuth
//	@Router			/accounts/delete [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	err := utils.ValidateToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GetToken(c)
	if err != nil {
		log.Println(err)
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	userID := claims["id"].(float64)

	if err := h.DB.Select("Task").Delete(&models.User{ID: uint64(userID)}).Error; err != nil {
		c.Status(http.StatusNoContent)
	}
}

// User profile godoc
//
//	@Summary		User profile
//	@Description	User profile
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.User
//	@in				header
//	@name			Authorization
//	@Security		BearerAuth
//	@Router			/accounts/profile [get]
func (h *Handler) Profile(c *gin.Context) {
	var user models.User
	err := utils.ValidateToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GetToken(c)
	if err != nil {
		log.Println(err)
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	userID := claims["id"]

	if err := h.DB.Where("id=?", userID).Preload("Task", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("id DESC")
	}).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
