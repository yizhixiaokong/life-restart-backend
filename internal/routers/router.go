package routers

import (
	"life-restart-backend/internal/api/expense"
	"life-restart-backend/internal/api/plan"
	"life-restart-backend/internal/api/readingentry"
	"life-restart-backend/internal/api/user"
	"life-restart-backend/internal/dao"
	expensesrv "life-restart-backend/internal/services/expense"
	plansrv "life-restart-backend/internal/services/plan"
	readingentrysrv "life-restart-backend/internal/services/readingentry"
	usersrv "life-restart-backend/internal/services/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userDAO := dao.NewUserDAO()
	userSrv := usersrv.NewUserService(userDAO)
	userAPI := user.NewUserHandler(userSrv)

	// 计划相关
	planDAO := dao.NewPlanDAO()
	planSrv := plansrv.NewPlanService(planDAO)
	planAPI := plan.NewPlanHandler(planSrv)

	// 开支相关
	expenseDAO := dao.NewExpenseDAO()
	expenseSrv := expensesrv.NewExpenseService(expenseDAO)
	expenseAPI := expense.NewExpenseHandler(expenseSrv)

	// 阅读条目相关
	readingEntryDAO := dao.NewReadingEntryDAO()
	readingEntrySrv := readingentrysrv.NewReadingEntryService(readingEntryDAO)
	readingEntryAPI := readingentry.NewReadingEntryHandler(readingEntrySrv)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userAPI.Register)
		v1.GET("/users/:id", userAPI.GetByID)
		v1.GET("/users", userAPI.GetAll)

		v1.GET("/plans", planAPI.GetAll)
		v1.GET("/plans/:id", planAPI.GetByID)
		v1.POST("/plans", planAPI.Create)
		v1.PUT("/plans/:id", planAPI.Update)
		v1.DELETE("/plans/:id", planAPI.Delete)

		v1.GET("/expenses", expenseAPI.GetAll)
		v1.GET("/expenses/:id", expenseAPI.GetByID)
		v1.POST("/expenses", expenseAPI.Create)
		v1.PUT("/expenses/:id", expenseAPI.Update)
		v1.DELETE("/expenses/:id", expenseAPI.Delete)

		v1.GET("/reading-entries", readingEntryAPI.GetAll)
		v1.GET("/reading-entries/:id", readingEntryAPI.GetByID)
		v1.POST("/reading-entries", readingEntryAPI.Create)
		v1.PUT("/reading-entries/:id", readingEntryAPI.Update)
		v1.DELETE("/reading-entries/:id", readingEntryAPI.Delete)
	}

	return router
}
