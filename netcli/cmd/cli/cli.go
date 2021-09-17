package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	noTargetErr = errors.New("no target input")
)

/*
cli official docs:
https://github.com/urfave/cli/blob/master/docs/v2/manual.md
- ns - Looks Up the NameServers for a Particular Host
- cname - Looks up the CNAME for a particular host
- ip - Looks up the IP addresses for a particular host
*/
func main() {
	app := &cli.App{
		Usage: "simple network related command tool - Query NameServers, CNAME, IPs",
		Commands: []*cli.Command{
			{
				Name:  "ns",
				Usage: "Looks Up the NameServers for a Particular Domain",
				Action: func(c *cli.Context) error {
					domain := c.Args().First()
					nsl, err := getNameServers(domain)
					if err != nil {
						return err
					}
					fmt.Println(nsl)
					return nil
				},
			},
			{
				Name:  "cname",
				Usage: "Looks up the CNAME for a particular fqdn",
				Action: func(c *cli.Context) error {
					fqdn := c.Args().First()
					cname, err := getCNAME(fqdn)
					if err != nil {
						return err
					}
					fmt.Println(cname)
					return nil
				},
			},
			{
				Name:  "ip",
				Usage: "Looks up the IP addresses for a particular fqdn",
				Action: func(c *cli.Context) error {
					fqdn := c.Args().First()
					ipl, err := getIP(fqdn)
					if err != nil {
						return err
					}
					fmt.Println(ipl)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getNameServers(host string) (string, error) {
	if host == "" {
		return "", noTargetErr
	}
	ns, err := net.LookupNS(host)
	if err != nil {
		return "", err
	}
	var nameServer []string
	for _, n := range ns {
		nameServer = append(nameServer, n.Host)
	}
	return strings.Join(nameServer, ", "), nil
}

func getCNAME(host string) (string, error) {
	if host == "" {
		return "", noTargetErr
	}
	cname, err := net.LookupCNAME(host)
	if err != nil {
		return "", err
	}
	return cname, nil
}

func getIP(host string) (string, error) {
	if host == "" {
		return "", noTargetErr
	}
	ips, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}
	var ipList []string
	for _, ip := range ips {
		ipList = append(ipList, ip.String())
	}
	return strings.Join(ipList, ", "), nil
}
