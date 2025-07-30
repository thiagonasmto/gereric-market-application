package routes

import (
	"gestao-vendas/controllers"
	"gestao-vendas/middlewares"
	"gestao-vendas/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	clientGroup := router.Group("/clients")
	{
		clientGroup.POST("/", controllers.CreateClient)
		clientGroup.GET("/", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.GetClients)
		clientGroup.GET("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.GetClient)
		clientGroup.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.UpdateClient)
		clientGroup.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.DeleteClient)
	}

	admGroup := router.Group("/adms")
	{
		admGroup.POST("/", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.CreateAdm)
		admGroup.GET("/", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.GetAdms)
		admGroup.GET("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.GetAdm)
		admGroup.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.UpdateAdm)
		admGroup.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.DeleteAdm)
	}

	productGroup := router.Group("/products")
	{
		productGroup.POST("/", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.CreateProduct)
		productGroup.GET("/", controllers.GetProducts)
		productGroup.GET("/:id", controllers.GetProduct)
		productGroup.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.UpdateProduct)
		productGroup.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.DeleteProduct)
	}

	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("/", middlewares.AuthMiddleware(), controllers.CreateOrder)
		orderGroup.GET("/", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"))
		orderGroup.GET("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.GetOrderById)
		orderGroup.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), controllers.UpdateOrder)
	}

	servicesGroup := router.Group("/services")
	{
		servicesGroup.GET("/generate-excel", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), services.GenerateExcelReport)
		servicesGroup.POST("/find-vogal", services.FindVogal)
		servicesGroup.GET("/rank-clients", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), services.GetRankActiveClients)
		servicesGroup.GET("/ordes-in-progress", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), services.GetOrdersInProgress)
		servicesGroup.GET("/summary", middlewares.AuthMiddleware(), middlewares.RequireRoleMiddleware("admin"), services.GetSummary)
	}

	router.POST("/login", services.Login)
}
