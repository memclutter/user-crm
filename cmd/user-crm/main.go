package main

import (
	"log"
	"os"

	usercrm "github.com/memclutter/user-crm"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "User CRM"
	app.Version = "0.0.1"
	app.Description = "CRM example. User database management."
	app.Action = usercrm.Start
	app.Flags = usercrm.FlagDefinitions

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
