package handler

import (
	"net/http"
	"strconv"
	timetracker "time-tracker"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary CreateUser
// @Tags user
// @Description create user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body timetracker.CreateUserInput true "user info" example({"name":"Иван","passportNumber":"1234 123456"})
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/ [post]
func (h *Handler) createUser(c *gin.Context) {
	var user timetracker.User
	if err := c.BindJSON(&user); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON for sign up")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := h.services.Authorisation.CreateUser(user)
	if err != nil {
		logrus.WithError(err).Error("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": id,
	}).Info("User created successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllUsersResponse struct {
	Data []timetracker.User `json:"data"`
}

// @Summary GetAllUsers
// @Tags user
// @Description Get all users
// @ID getAllUsers-user
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/ [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	userList, err := h.services.Authorisation.GetAllUsers()
	if err != nil {
		logrus.WithError(err).Error("Failed to get all users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all users"})
		return
	}

	logrus.Info("Retrieved all users successfully")

	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: userList,
	})
}

// @Summary UpdateUser
// @Tags user
// @Description update user
// @ID update-user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param input body timetracker.UpdateUserInput true "User info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid user ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var input timetracker.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON for update user")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = h.services.Authorisation.UpdateUser(id, input)
	if err != nil {
		logrus.WithError(err).Error("Failed to update user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": id,
	}).Info("User updated successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type statusResponse struct {
	Status string `json:"status"`
}

// @Summary DeleteUser
// @Tags user
// @Description delete user
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid user ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.services.Authorisation.DeleteUser(id)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": id,
	}).Info("User deleted successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary SearchUser
// @Tags user
// @Description search user
// @ID search-user
// @Accept  json
// @Produce  json
// @Param name query string false "Name filter" example(Иван)
// @Param passportNumber query string false "Passport number filter" example(1234 123456)
// @Param page query int false "Page number" example(1)
// @Param pageSize query int false "Page size" example(10)
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/search [get]
func (h *Handler) searchUsers(c *gin.Context) {
	// Извлекаем параметры фильтрации из строки запроса
	nameFilter := c.Query("name")
	passportNumberFilter := c.Query("passportNumber")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1 // По умолчанию первая страница, если параметр не задан или не удалось преобразовать
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 10 // По умолчанию размер страницы 10 элементов, если параметр не задан или не удалось преобразовать
	}

	logrus.WithFields(logrus.Fields{
		"nameFilter":           nameFilter,
		"passportNumberFilter": passportNumberFilter,
		"page":                 page,
		"pageSize":             pageSize,
	}).Debug("Searching users with filters")

	// Вызываем сервис или репозиторий для выполнения поиска пользователей
	users, err := h.services.Authorisation.Search(nameFilter, passportNumberFilter, page, pageSize)
	if err != nil {
		logrus.WithError(err).Error("Failed to search users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"count": len(users),
	}).Info("Users search completed successfully")

	// Возвращаем найденных пользователей в виде JSON
	c.JSON(http.StatusOK, gin.H{"users": users})
}
