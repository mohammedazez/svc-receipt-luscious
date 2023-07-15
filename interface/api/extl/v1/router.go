package v1

import (
	"fmt"
	"log"
	"svc-receipt-luscious/utils/config/mysql"

	"github.com/labstack/echo/v4"
)

func API(route *echo.Echo) {

	// instance mysql
	mysqlDB, err := mysql.GetMysqlDB()
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(mysqlDB)

}
