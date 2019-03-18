package main

import (
	"net/http"

	"github.com/sereiner/higo/middleware"

	"github.com/sereiner/higo"
)

func main() {
	// Setup
	e := higo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c higo.Context) error {
		c.Logger().Info("ok")
		return c.JSON(http.StatusOK, "OK")
	})
	e.Start(":1323")
	// Start server
	// go func() {
	// 	if err := e.Start(":1323"); err != nil {
	// 		e.Logger.Info("shutting down the server")
	// 	}
	// }()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// time.Sleep(time.Second * 2)
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// if err := e.Shutdown(ctx); err != nil {
	// 	e.Logger.Fatal(err)
	// }
}
