package usercrm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/memclutter/user-crm/models"
	"github.com/urfave/cli"
)

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

	return nil
}
