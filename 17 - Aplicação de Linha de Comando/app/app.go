package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	// app.Commands = []cli.Command{
	// 	{
	// 		Name:  "ip",
	// 		Usage: "Busca IPS de endereços na internet",
	// 		Flags: []cli.Flag{
	// 			cli.StringFlag{
	// 				Name:  "host",
	// 				Value: "devbook.con.br",
	// 			},
	// 		},
	// 		// Action: func(c *cli.Context) {
	// 		// },
	// 		Action: buscaIps,
	// 	},
	// }

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.con.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet",
			Flags:  flags,
			Action: buscaIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscaIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ns := range servidores {
		// fmt.Println(ns)
		fmt.Println(ns.Host)
	}
}
