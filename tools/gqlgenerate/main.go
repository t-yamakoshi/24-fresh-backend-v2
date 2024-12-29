package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

func main() {
	os.Exit(run())
}

func run() (code int) {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		log.Printf("failed to load config. err = %+v", err)
		return 1
	}

	plugin := &modelgen.Plugin{
	}
	if err := api.Generate(cfg, api.ReplacePlugin(plugin)); err != nil {
		log.Printf("failed to generate code. err = %+v", err)
		return 1
	}
	return 0
}
