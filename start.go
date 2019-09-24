package usercrm

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/memclutter/user-crm/models"
	"github.com/urfave/cli"
)

// @title User CRM
// @version 1.0
// @description CRM example. User database management.
//
// @host localhost:9000
// @BasePath /api/
func Start(c *cli.Context) error {

	// Parse flags
	flags := NewFlags(c)

	// Open database connection
	db, err := gorm.Open("postgres", flags.PsqlURL)
	if err != nil {
		return fmt.Errorf("error open database connection: %s", err)
	}
	defer db.Close()

	// Debug mode
	db.LogMode(flags.Debug)

	// Apply migrations
	db.AutoMigrate(&models.User{}, &models.Country{})

	// Init http router
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Debug = flags.Debug

	// Create custom validator
	customValidator, err := NewCustomValidator()
	if err != nil {
		return err
	}

	// Set dependencies
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			c.Set("flags", flags)
			c.Set("validator", customValidator)
			return next(c)
		}
	})

	api := e.Group("/api")

	// Swagger docs
	api.Static("/docs", flags.DocsRoot)

	// Routes
	NewCountries(api)
	NewUsers(api)

	// Run server
	go func() {
		if err := e.Start(":9000"); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	// Wait 10 sec for server shutdown
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}
