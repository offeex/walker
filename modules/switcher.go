package modules

import (
	"context"

	"github.com/abenz1267/walker/config"
)

type Switcher struct {
	general config.GeneralModule
	cfg     *config.Config
}

func (s Switcher) Placeholder() string {
	if s.general.Placeholder == "" {
		return "switcher"
	}

	return s.general.Placeholder
}

func (s Switcher) SwitcherOnly() bool {
	return false
}

func (s Switcher) Entries(ctx context.Context, term string) []Entry {
	entries := []Entry{}

	for _, v := range s.cfg.Enabled {
		if v == "switcher" {
			continue
		}

		e := Entry{
			Label:      v,
			Sub:        "switcher",
			Exec:       "",
			Categories: []string{"switcher"},
			Class:      "switcher",
			Matching:   Fuzzy,
		}

		entries = append(entries, e)
	}

	return entries
}

func (s Switcher) Prefix() string {
	return s.general.Prefix
}

func (s Switcher) Name() string {
	return "switcher"
}

func (s *Switcher) Setup(cfg *config.Config) {
	s.general.Prefix = cfg.Builtins.Switcher.Prefix
	s.cfg = cfg

	s.general.IsSetup = true
}

func (s *Switcher) SetupData(cfg *config.Config) {}

func (s Switcher) IsSetup() bool {
	return s.general.IsSetup
}

func (s Switcher) Refresh() {}
