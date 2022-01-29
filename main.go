package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2" // imports as package "cli"
)

func main() {
	fmt.Println("I am GoLang script!")

	cli_example()
}

func cli_example() {
	app := &cli.App{
		Name:  "GoLang script",
		Usage: "CLI example",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "Foo",
				Usage:   "Your name",
			},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			fmt.Printf("Hello %q, welcome to the world of CLI.", name)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
