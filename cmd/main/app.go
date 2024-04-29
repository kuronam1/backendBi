package main

import (
	"flag"
	"log"
	"sbitnev_back/internal/app"
	"sbitnev_back/internal/config"
)

func main() {
	var configPath string
	flag.StringVar(
		&configPath, "cfg",
		"./internal/config/config.yaml",
		"path to config file",
	)
	flag.Parse()

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
