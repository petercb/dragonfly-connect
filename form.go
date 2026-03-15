package main

import (
	"log/slog"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
)

// ServerListForm handles the menu form submission.
type ServerListForm struct {
	Log     *slog.Logger
	Servers []ServerConfig
}

// Submit is called when the user presses one of the server buttons in the form.
func (f ServerListForm) Submit(submitter form.Submitter, pressed form.Button, tx *world.Tx) {
	if p, ok := submitter.(*player.Player); ok {
		for _, s := range f.Servers {
			if s.Name == pressed.Text {
				f.Log.Info("Transferring player", "player", p.Name(), "target", s.Name, "address", s.Address)
				if err := p.Transfer(s.Address); err != nil {
					f.Log.Error("Transfer failed", "player", p.Name(), "err", err)
				}
				return
			}
		}
	}
}

// CreateServerMenu dynamically creates the form with the available servers from config.
func CreateServerMenu(log *slog.Logger, servers []ServerConfig) form.Menu {
	f := ServerListForm{
		Log:     log,
		Servers: servers,
	}
	m := form.NewMenu(f, "Bedrock Connect").WithBody("Select a server to join:")
	for _, s := range servers {
		m = m.WithButtons(form.NewButton(s.Name, s.Image))
	}
	return m
}
