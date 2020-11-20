package main

import (
	"log"
	"strings"

	"github.com/dawidd6/go-spotify-dbus"
	"github.com/godbus/dbus"
)

type Config struct {
	SkipKeywords []string `json:"skip_keywords"`
}

func defaultConfig() *Config {
	return &Config{
		SkipKeywords: []string{"instrument"},
	}
}

type CLI struct {
	config    *Config
	conn      *dbus.Conn
	listeners *spotify.Listeners
}

func NewCLI(config *Config) (*CLI, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	listeners := spotify.NewListeners()

	cli := &CLI{
		config:    config,
		conn:      conn,
		listeners: listeners,
	}

	cli.init()
	return cli, nil
}

func (c *CLI) Run() {
	spotify.Listen(c.conn, c.listeners)
}

func (c *CLI) init() {
	c.listeners.OnMetadata = c.onMetadata
	c.listeners.OnPlaybackStatus = c.onPlaybackStatus
	c.listeners.OnServiceStart = c.onServiceStart
	c.listeners.OnServiceStop = c.onServiceStop
	c.listeners.OnError = c.onError
}

func contains(str string, substrs []string) bool {
	for _, substr := range substrs {
		if strings.Contains(str, substr) {
			return true
		}
	}
	return false
}

func (c *CLI) onMetadata(metadata *spotify.Metadata) {
	log.Printf("[playing] %s\n", metadata.Title)

	if title := strings.ToLower(metadata.Title); contains(title, c.config.SkipKeywords) {
		if err := spotify.SendNext(c.conn); err != nil {
			log.Printf("[error] %v\n", err)
		} else {
			log.Printf("[skipped] %s\n", title)
		}
	}
}

func (c *CLI) onPlaybackStatus(status spotify.PlaybackStatus) {
	log.Printf("[status] %s", status)
}

func (c *CLI) onServiceStart() {
	log.Println("[service] Started")
}

func (c *CLI) onServiceStop() {
	log.Println("[service] Stopped")
}

func (c *CLI) onError(err error) {
	log.Printf("[error] %v", err)
}
