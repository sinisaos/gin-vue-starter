package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sinisaos/gin-vue-starter/pkg/models"
	"gorm.io/gorm"
)

const DefaultPageSize = 15

// AllTasks godoc
//
//	@Summary		List tasks
//	@Description	Show all tasks.
//	@Tags			Task
//	@Produce		json
//	@Param			page		query	string	false	"page"
//	@Param			page_size	query	string	false	"page_size"
//	@Success		200			{array}	[]models.Task
//	@Router			/tasks [get]
func (h *Handler) AllTasks(c *gin.Context) {
	pageQuery := c.Query("page")
	sizeQuery := c.Query("page_size")

	var tasks []models.Task
	// pagination
	if pageQuery == "" {
		pageQuery = "1"
	}
	page, _ := strconv.Atoi(pageQuery)
	size, _ := strconv.Atoi(sizeQuery)

	perPage := size
	offset := perPage * (page - 1)

	if pageQuery != "" && sizeQuery != "" {
		h.DB.Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "user_name")
		}).Order("id desc").Limit(perPage).Offset(offset).Find(&tasks)
	} else {
		h.DB.Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, user_name")
		}).Order("id desc").Limit(DefaultPageSize).Find(&tasks)
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks, "page": page})
}

// SingleTask godoc
//
//	@Summary		Single Task
//	@Description	Show single Task.
//	@Tags			Task
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	models.Task
//	@Failure		404	{string}	string	"Record not found!"
//	@Router			/tasks/{id} [get]
func (h *Handler) SingleTask(c *gin.Context) {
	ID := c.Param("id")
	var task models.Task
	if err := h.DB.Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "user_name")
	}).Where("id = ?", ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// CreateTask godoc
//
//	@Summary		Create Task
//	@Description	Create a new Task.
//	@Tags			Task
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.TaskValidate	true	"body"
//	@Success		200		{object}	models.Task
//	@Failure		404		{string}	string	err.Error()
//	@in				header
//	@name			Authorization
//	@Security		BearerAuth
//	@Router			/tasks [post]
func (h *Handler) CreateTask(c *gin.Context) {
	var taskData models.TaskValidate
	if err := c.ShouldBindJSON(&taskData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Name:      taskData.Name,
		Completed: taskData.Completed,
		UserID:    taskData.UserID,
	}
	h.DB.Create(&task)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// UpdateTask godoc
//
//	@Summary		Update Task
//	@Description	Update single Task.
//	@Tags			Task
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"id"
//	@Param			body	body		models.TaskValidate	true	"body"
//	@Success		200		{object}	models.Task
//	@Failure		404		{string}	string	"Record not found!"
//	@in				header
//	@name			Authorization
//	@Security		BearerAuth
//	@Router			/tasks/{id} [patch]
func (h *Handler) UpdateTask(c *gin.Context) {
	ID := c.Param("id")
	var task models.Task
	if err := h.DB.Where("id = ?", ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var taskData models.TaskValidate
	if err := c.ShouldBindJSON(&taskData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Model(&task).Where("id = ?", ID).Updates(models.Task{
		Name:      taskData.Name,
		Completed: taskData.Completed,
		UserID:    taskData.UserID,
	})
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DeleteTask godoc
//
//	@Summary		Delete Task
//	@Description	Delete single Task.
//	@Tags			Task
//	@Param			id	path	string	true	"id"
//	@Success		204
//	@Failure		404	{string}	string	"Record not found!"
//	@in				header
//	@name			Authorization
//	@Security		BearerAuth
//	@Router			/tasks/{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	ID := c.Param("id")
	var task models.Task
	if err := h.DB.Where("id = ?", ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	h.DB.Delete(&task)
	c.Status(http.StatusNoContent)
}
