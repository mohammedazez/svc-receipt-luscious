package v1

import (
	"log"
	categoryService "svc-receipt-luscious/core/category"
	healthCheckService "svc-receipt-luscious/core/healthCheck"
	ingredientService "svc-receipt-luscious/core/ingredient"
	categoryRepositoryMysql "svc-receipt-luscious/infrastructure/repository/mysql/category"
	healtCheckRepositoryMysql "svc-receipt-luscious/infrastructure/repository/mysql/healthCheck"
	ingredientRepositoryMysql "svc-receipt-luscious/infrastructure/repository/mysql/ingredient"
	categoryHandlerV1 "svc-receipt-luscious/interface/api/extl/v1/category/handler"
	healthCheckHandlerV1 "svc-receipt-luscious/interface/api/extl/v1/healthCheck"
	ingredientHandlerV1 "svc-receipt-luscious/interface/api/extl/v1/ingredient/handler"
	"svc-receipt-luscious/utils/config/mysql"

	"github.com/labstack/echo/v4"
)

func API(route *echo.Echo) {

	// instance mysql
	mysqlDB, err := mysql.GetMysqlDB()
	if err != nil {
		log.Panic(err.Error())
	}

	// instance repo
	healthCheckRepositoryMysql := healtCheckRepositoryMysql.NewRepository(mysqlDB)
	ingredientRepositoryMysql := ingredientRepositoryMysql.NewRepository(mysqlDB)
	categoryRepositoryMysql := categoryRepositoryMysql.NewRepository(mysqlDB)

	// instance service
	healthCheckService := healthCheckService.NewService(healthCheckRepositoryMysql)
	ingredientService := ingredientService.NewService(ingredientRepositoryMysql)
	categoryService := categoryService.NewService(categoryRepositoryMysql)

	// instance handler
	healthCheckHandlerV1 := healthCheckHandlerV1.NewHandler(healthCheckService)
	ingredientHandlerV1 := ingredientHandlerV1.NewHandler(ingredientService)
	categoryHandlerV1 := categoryHandlerV1.NewHandler(categoryService)

	// V1 routes
	v1Route := route.Group("/v1")

	healthCheckRouteV1 := v1Route.Group("/health")
	healthCheckRouteV1.GET("", healthCheckHandlerV1.Get)

	// ingredient
	ingredientRouteV1 := v1Route.Group("/ingredient")
	ingredientRouteV1.GET("", ingredientHandlerV1.List)
	ingredientRouteV1.POST("", ingredientHandlerV1.Insert)
	ingredientRouteV1.PUT("/:ingredient_id", ingredientHandlerV1.Update)
	ingredientRouteV1.DELETE("/:ingredient_id", ingredientHandlerV1.Delete)

	// category
	categoryRouteV1 := v1Route.Group("/category")
	categoryRouteV1.GET("", categoryHandlerV1.List)
	// categoryRouteV1.GET("/:category_id", categoryHandlerV1.Detail)
	// categoryRouteV1.POST("", categoryHandlerV1.Insert)
	// categoryRouteV1.PUT("/:category_id", categoryHandlerV1.Update)
	// categoryRouteV1.DELETE("/:category_id", categoryHandlerV1.Delete)

	// recipe
	// recipeRouteV1 := v1Route.Group("/recipe")
	// recipeRouteV1.GET()
	// recipeRouteV1.GET()
	// recipeRouteV1.POST()
	// recipeRouteV1.PUT()
	// recipeRouteV1.DELETE()

}
