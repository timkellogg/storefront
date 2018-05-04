package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config - represents options present in config.toml
type Config struct {
	Address string
	Version string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
