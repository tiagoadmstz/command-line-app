package app

import "github.com/urfave/cli"

// Generate return the command line application ready to start
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Searchs for hostnames and IPs on the internet"

	return app
}
