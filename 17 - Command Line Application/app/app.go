package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will be return the command line application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IP'S and Server Names on Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP's on the internet",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "servers",
			Usage:  "Search server name on the Internet",
			Flags:  flags,
			Action: searchServers,
		},
	}
	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro_r := net.LookupIP(host)

	if erro_r != nil {
		log.Fatal(erro_r)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, erro_r := net.LookupNS(host) // name server

	if erro_r != nil {
		log.Fatal(erro_r)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
