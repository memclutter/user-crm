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
		Name:     "docsRoot",
		Usage:    "Root folder with api docs",
		EnvVar:   "DOCS_ROOT",
		Value:    "docs",
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
	Debug    bool
	DocsRoot string
	PsqlURL  string
}

func NewFlags(c *cli.Context) *Flags {
	return &Flags{
		Debug:    c.Bool("debug"),
		DocsRoot: c.String("docsRoot"),
		PsqlURL:  c.String("psqlURL"),
	}
}
