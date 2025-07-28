package routes

import (
	"golang_starter_kit_2025/app/controllers"
	"golang_starter_kit_2025/app/features/category"
	"golang_starter_kit_2025/app/features/product"
	"golang_starter_kit_2025/app/middleware"
	"golang_starter_kit_2025/app/services"
	"golang_starter_kit_2025/facades"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	controller := controllers.Controller{}
	route.GET("", controller.HelloWorld)

	authService := services.AuthService{}
	authController := controllers.NewAuthController(authService)
	route.PUT("/auth/login", authController.Login)
	authRoutes := route.Group("/auth").Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/logout", authController.Logout)
		authRoutes.GET("/refresh", authController.Refresh)
	}

	categoryController := category.NewCategoryController()
	categoryRoutes := route.Group("/categories")
	{
		categoryRoutes.GET("/", categoryController.List)
		categoryRoutes.GET("/:id", categoryController.Get)
		categoryRoutes.PUT("/", categoryController.Put)
		categoryRoutes.DELETE("/:id", categoryController.Delete)
	}

	productController := product.NewProductController()
	productRoutes := route.Group("/products")
	{
		productRoutes.GET("/", productController.GetAll)
		productRoutes.GET("/:id", productController.GetByID)
		productRoutes.PUT("/", productController.Put)
		productRoutes.DELETE("/:id", productController.Delete)
	}

	userController := controllers.NewUserController()
	userRoutes := route.Group("/users", middleware.AuthMiddleware())
	{
		userRoutes.GET("", userController.List)
		userRoutes.GET("/:id", userController.Get)
		userRoutes.PUT("", userController.Put)
		userRoutes.DELETE("/:id", userController.Delete)
		userRoutes.POST("/:id/roles", userController.AssignRoles)
		userRoutes.GET("/:id/roles", userController.GetRoles)
	}

	roleService := services.RoleService{}
	roleController := controllers.NewRoleController(roleService)
	roleRoutes := route.Group("/roles", middleware.AuthMiddleware())
	{
		roleRoutes.GET("", roleController.List)
		roleRoutes.PUT("", roleController.Put)
		roleRoutes.DELETE("/:id", roleController.Delete)
		roleRoutes.POST("/:id/permissions", roleController.AssignPermissions)
		roleRoutes.GET("/:id/permissions", roleController.GetPermissions)
	}

	permissionService := services.PermissionService{}
	permissionController := controllers.NewPermissionController(permissionService)
	permissionRoutes := route.Group("/permissions", middleware.AuthMiddleware())
	{
		permissionRoutes.GET("", permissionController.List)
		permissionRoutes.PUT("", permissionController.Put)
		permissionRoutes.DELETE("/:id", permissionController.Delete)
	}

	fileController := controllers.NewFileController()
	fileRoutes := route.Group("/file")
	{
		fileRoutes.GET("/:key/:filename", fileController.ServeFile)
	}

	databaseController := controllers.NewDatabaseController()
	databaseRoutes := route.Group("/api/database")
	{
		databaseRoutes.GET("/status", databaseController.GetConnectionStatus)
		databaseRoutes.GET("/health", databaseController.HealthCheck)
		databaseRoutes.GET("/test", databaseController.TestConnection)
	}

	route.GET("/health", func(c *gin.Context) {
		sqlDB, err := facades.DB.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get facades connection",
				"error":   err.Error(),
			})
			return
		}

		err = sqlDB.Ping()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "facades connection failed",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "facades is connected",
			"facades": "supply_chain_retail",
		})
	})

	route.GET("/health/databases", func(c *gin.Context) {
		manager := facades.GetManager()
		health := make(map[string]interface{})
		connections := []string{"mysql", "postgres", "mysql_secondary"}

		allHealthy := true
		for _, connName := range connections {
			if manager.IsConnected(connName) {
				stats, err := manager.GetConnectionStats(connName)
				if err == nil {
					health[connName] = gin.H{
						"status": "healthy",
						"stats":  stats,
					}
				} else {
					health[connName] = gin.H{
						"status": "unhealthy",
						"error":  err.Error(),
					}
					allHealthy = false
				}
			} else {
				health[connName] = gin.H{
					"status": "disconnected",
				}
				allHealthy = false
			}
		}

		statusCode := http.StatusOK
		if !allHealthy {
			statusCode = http.StatusServiceUnavailable
		}

		c.JSON(statusCode, gin.H{
			"overall_health": allHealthy,
			"connections":    health,
		})
	})
}
