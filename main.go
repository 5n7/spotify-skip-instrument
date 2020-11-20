package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func configPath() (string, error) {
	cfg, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cfg, "spotify-skip-instrument", "config.json"), err
}

func loadConfig(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func run() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	var config *Config
	if _, err := os.Stat(path); os.IsNotExist(err) {
		config = defaultConfig()
	} else {
		cfg, err := loadConfig(path)
		if err != nil {
			return err
		}
		config = cfg
	}

	cli, err := NewCLI(config)
	if err != nil {
		return err
	}

	cli.Run()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
