package state

import (
	"time"

	"github.com/abenz1267/walker/config"
	"github.com/abenz1267/walker/modules"
	"github.com/abenz1267/walker/modules/clipboard"
)

type AppState struct {
	Started             time.Time
	IsMeasured          bool
	IsService           bool
	IsRunning           bool
	HasUI               bool
	Clipboard           modules.Workable
	Dmenu               modules.Workable
	Password            bool
	ExplicitModules     []string
	ExplicitConfig      string
	ExplicitStyle       string
	ExplicitPlaceholder string
	KeepSort            bool
}

func Get() *AppState {
	return &AppState{
		Started:        time.Now(),
		IsService:      false,
		IsRunning:      false,
		IsMeasured:     false,
		HasUI:          false,
		ExplicitConfig: "config.json",
		ExplicitStyle:  "style.css",
	}
}

func (app *AppState) StartServiceableModules(cfg *config.Config) {
	app.Clipboard = &clipboard.Clipboard{}
}
