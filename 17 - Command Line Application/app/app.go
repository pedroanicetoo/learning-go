package app

import "github.com/urfave/cli"

// Generate will be return the command line application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IP'S and Server Names on Internet"

	return app
}
