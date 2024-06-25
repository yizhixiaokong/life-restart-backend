package routers

import (
	"life-restart-backend/internal/api/user"
	"life-restart-backend/internal/dao"
	usersrv "life-restart-backend/internal/services/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userDAO := dao.NewUserDAO()
	userSrv := usersrv.NewUserService(userDAO)
	userapi := user.NewUserHandler(userSrv)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userapi.Register)
		v1.GET("/users/:id", userapi.GetByID)
	}

	return router
}
