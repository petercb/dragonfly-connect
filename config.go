package main

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Image   string `json:"image"`
}

type Config struct {
	Servers []ServerConfig `json:"servers"`
}

func loadConfig(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
