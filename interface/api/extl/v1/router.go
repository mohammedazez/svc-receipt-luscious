package v1

import (
	"log"
	categoryService "svc-receipt-luscious/core/category"
	healthCheckService "svc-receipt-luscious/core/healthCheck"
	ingredientService "svc-receipt-luscious/core/ingredient"
	recipeService "svc-receipt-luscious/core/recipe"
	categoryRepositoryMysql "svc-receipt-luscious/infrastructure/repository/mysql/category"
	healtCheckRepositoryMysql "svc-receipt-luscious/infrastructure/repository/mysql/healthCheck"
	ingredientRepositoryMysql "svc-receipt-luscious/infrastructure/repository/mysql/ingredient"
	recipeRepositoryMysql "svc-receipt-luscious/infrastructure/repository/mysql/recipe"
	categoryHandlerV1 "svc-receipt-luscious/interface/api/extl/v1/category/handler"
	healthCheckHandlerV1 "svc-receipt-luscious/interface/api/extl/v1/healthCheck"
	ingredientHandlerV1 "svc-receipt-luscious/interface/api/extl/v1/ingredient/handler"
	recipeHandlerV1 "svc-receipt-luscious/interface/api/extl/v1/recipe/handler"
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
	recipeRepositoryMysql := recipeRepositoryMysql.NewRepository(mysqlDB)

	// instance service
	healthCheckService := healthCheckService.NewService(healthCheckRepositoryMysql)
	ingredientService := ingredientService.NewService(ingredientRepositoryMysql)
	categoryService := categoryService.NewService(categoryRepositoryMysql)
	recipeService := recipeService.NewService(recipeRepositoryMysql)

	// instance handler
	healthCheckHandlerV1 := healthCheckHandlerV1.NewHandler(healthCheckService)
	ingredientHandlerV1 := ingredientHandlerV1.NewHandler(ingredientService)
	categoryHandlerV1 := categoryHandlerV1.NewHandler(categoryService)
	recipeHandlerV1 := recipeHandlerV1.NewHandler(recipeService)

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
	categoryRouteV1.POST("", categoryHandlerV1.Insert)
	categoryRouteV1.PUT("/:category_id", categoryHandlerV1.Update)
	categoryRouteV1.DELETE("/:category_id", categoryHandlerV1.Delete)

	// recipe
	recipeRouteV1 := v1Route.Group("/recipe")
	recipeRouteV1.GET("", recipeHandlerV1.List)
	// recipeRouteV1.GET("/:recipe_id", recipeHandlerV1.Detail)
	recipeRouteV1.POST("", recipeHandlerV1.Insert)
	recipeRouteV1.PUT("/:recipe_id", recipeHandlerV1.Update)
	recipeRouteV1.DELETE("/:recipe_id", recipeHandlerV1.Delete)

}
