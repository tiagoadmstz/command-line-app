package app

import (
	"command-line-app/util"
	"net"

	"github.com/urfave/cli"
)

// Generate return the command line application ready to start
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Searchs for hostnames and IPs on the internet"
	configure(app)
	return app
}

func configure(app *cli.App) {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IP adresses on the internet",
			Flags:  flags,
			Action: searchIp,
		},
		{
			Name:   "servers",
			Usage:  "Search for servers names on the internet",
			Flags:  flags,
			Action: searchServers,
		},
	}
}

func searchIp(context *cli.Context) {
	util.Search(context, func(host string) ([]net.IP, error) {
		return net.LookupIP(host)
	})
}

func searchServers(context *cli.Context) {
	util.Search(context, func(host string) ([]string, error) {
		servers, err := net.LookupNS(host)
		serversStr := util.Map(servers, func(server *net.NS) string {
			return server.Host
		})
		return serversStr, err
	})
}
