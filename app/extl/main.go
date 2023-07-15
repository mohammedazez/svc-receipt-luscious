package main

import (
	"context"
	"os"
	"os/signal"
	routes "svc-receipt-luscious/interface/api/extl/v1"
	"svc-receipt-luscious/utils/config"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	config.LoadConfig()
	e := echo.New()

	// route
	routes.API(e)

	// Start server
	go func() {
		if err := e.Start(":" + os.Getenv("APP_PORT_LUSCIOUS")); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
