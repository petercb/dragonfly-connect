package main

import (
	"log/slog"
	"os"

	"github.com/df-mc/dragonfly/server"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	cfg, err := loadConfig("servers.json")
	if err != nil {
		log.Error("Failed to load servers.json. Make sure the file exists and is valid JSON.", "err", err)
		return
	}
	log.Info("Loaded target servers", "count", len(cfg.Servers))

	conf, err := server.DefaultConfig().Config(log)
	if err != nil {
		log.Error("Error determining dragonfly server config", "err", err)
		return
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.Listen()

	log.Info("BedrockConnect ready! Connect to the proxy via Bedrock client.")

	for p := range srv.Accept() {
		log.Info("Player joined", "name", p.Name())
		p.SendForm(CreateServerMenu(log, cfg.Servers))
	}
}
