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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IP's on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchIPs,
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
