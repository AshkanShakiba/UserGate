package main

import (
	"database/sql"
	"fmt"
	"github.com/AshkanShakiba/UserGate/internal/config"
	"github.com/AshkanShakiba/UserGate/pkg/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	conf, err := config.Configure()
	if err != nil {
		log.Fatalf("loading configuration failed: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	s := server.NewServer(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/user", func(c echo.Context) error {
		return handleServerMethod(c, s.GetUser)
	})
	e.POST("/user", func(c echo.Context) error {
		return handleServerMethod(c, s.CreateUser)
	})

	port := fmt.Sprintf(":%s", conf.Port)
	log.Printf("Starting server on %s", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// handleServerMethod wraps the server's HTTP handlers to capture panics or errors
func handleServerMethod(c echo.Context, handler func(http.ResponseWriter, *http.Request)) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Internal Server Error",
			})
		}
	}()

	w := c.Response().Writer
	req := c.Request()

	handler(w, req)

	if c.Response().Status == 0 {
		c.Response().WriteHeader(http.StatusOK)
	}

	return nil
}
