package usercrm

import "github.com/urfave/cli"

var FlagDefinitions = []cli.Flag{
	cli.BoolFlag{
		Name:     "debug",
		Usage:    "Enable debug mode",
		EnvVar:   "DEBUG",
		Required: false,
	},
	cli.StringFlag{
		Name:     "psqlURL",
		Usage:    "PostgresSQL database connection string. Example 'host=<host> user=<user> sslmode=disable password=<password> dbname=<dbname>' or 'postgres://<user>:<password>@<host>:<port>/<dbname>?sslmode=disable",
		EnvVar:   "PSQL_URL",
		Required: true,
	},
}

type Flags struct {
	Debug   bool
	PsqlURL string
}

func NewFlags(c *cli.Context) *Flags {
	return &Flags{
		Debug:   c.Bool("debug"),
		PsqlURL: c.String("psqlURL"),
	}
}
