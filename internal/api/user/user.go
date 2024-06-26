package user

import (
	"life-restart-backend/internal/dao/models"
	userservice "life-restart-backend/internal/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	userSrv *userservice.UserService
}

func NewUserHandler(userSrv *userservice.UserService) *Handler {
	return &Handler{
		userSrv: userSrv,
	}
}

func (api *Handler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := api.userSrv.RegisterUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "id": id})
}

func (api *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := api.userSrv.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (api *Handler) GetAll(c *gin.Context) {
	users, err := api.userSrv.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
