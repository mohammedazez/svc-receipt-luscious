package config

import (
	"svc-receipt-luscious/utils/config/mysql"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	godotenv.Load()

	mysql.Init()
}
