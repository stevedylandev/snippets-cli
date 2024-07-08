package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type ResponseData struct {
	IpfsHash    string `json:"IpfsHash"`
	PinSize     int    `json:"PinSize"`
	Timestamp   string `json:"Timestamp"`
	IsDuplicate bool   `json:"isDuplicate"`
}

type UploadPayload struct {
	Content string `json:"content"`
	Name    string `json:"name"`
	Lang    string `json:"lang"`
}

func main() {
	app := &cli.App{
		Name:  "snip",
		Usage: "A CLI for snippets.so, a clean and simple code sharing tool.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "nil",
				Usage:   "Custom name for snippet. If none provided it will default to the name of the file.",
			},
		},
		Action: func(ctx *cli.Context) error {
			filePath := ctx.Args().First()
			name := ctx.String("name")
			if filePath == "" {
				return errors.New("no file path provided")
			}
			_, err := UploadSnip(filePath, name)
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
