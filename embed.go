package checkingautomata

import (
	_ "embed"
)

// DefaultConfig holds the default configuration.
//
//go:embed config/config.default.yaml
var DefaultConfig []byte
