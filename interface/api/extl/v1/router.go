package v1

import (
	"log"
	healthCheckService "svc-receipt-luscious/core/healthCheck"
	ingredientService "svc-receipt-luscious/core/ingredient"
	healtCheckRepositoryMysql "svc-receipt-luscious/infrastructure/repository/mysql/healthCheck"
	ingredientRepositoryMysql "svc-receipt-luscious/infrastructure/repository/mysql/ingredient"
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

	// instance service
	healthCheckService := healthCheckService.NewService(healthCheckRepositoryMysql)
	ingredientService := ingredientService.NewService(ingredientRepositoryMysql)

	// instance handler
	healthCheckHandlerV1 := healthCheckHandlerV1.NewHandler(healthCheckService)
	ingredientHandlerV1 := ingredientHandlerV1.NewHandler(ingredientService)

	// V1 routes
	v1Route := route.Group("/v1")

	healthCheckRouteV1 := v1Route.Group("/health")
	healthCheckRouteV1.GET("", healthCheckHandlerV1.Get)

	ingredientRouteV1 := v1Route.Group("/ingredient")
	ingredientRouteV1.GET("", ingredientHandlerV1.List)

}
