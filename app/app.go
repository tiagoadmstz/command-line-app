package app

import (
	"fmt"
	"log"
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
	host := context.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(context *cli.Context) {
	host := context.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
