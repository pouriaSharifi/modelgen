package main

import (
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"log"
	"modelgen/mgenV2"

	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{Commands: []*cli.Command{
		{
			Name:  "generate",
			Usage: "generate graphql schema",
			Action: func(c *cli.Context) error {
				gqlgenConf, err := config.LoadConfigFromDefaultLocations()
				if err != nil {
					log.Panicln("failed to load config", err.Error())

					os.Exit(2)
				}

				log.Println("generating schema...")
				err = api.Generate(gqlgenConf, api.ReplacePlugin(mgenV2.New()))
				log.Println("schema generation done...")

				if err != nil {
					fmt.Println(err)
					log.Panicln(err.Error())
					os.Exit(3)
				}
				os.Exit(0)
				return nil
			},
		},
	}}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
