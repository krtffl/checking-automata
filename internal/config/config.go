package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	checkingautomata "github.com/krtffl/checking-automata"
)

type Config struct {
	Mailgun Mailgun `mapstructure:"mailgun" yaml:"mailgun"`
	Browser Browser `mapstructure:"browser" yaml:"browser"`
}

// Mailgun holds the configuration for mailgun
type Mailgun struct {
	// Whether email service is enabled
	Enable bool `mapstructure:"enable" yaml:"enable"`

	// Domain from which mails will be sent
	Domain string `mapstructure:"domain" yaml:"domain"`

	// API key
	Key string `mapstructure:"key" yaml:"key"`

	From    string `mapstructure:"from"    yaml:"from"`
	To      string `mapstructure:"to"      yaml:"to"`
	Subject string `mapstructure:"subject" yaml:"subject"`
}

type Browser struct {
	// wsURL to attach to
	Address string `mapstructure:"address" yaml:"address"`

	Timeout int `mapstructure:"timeout" yaml:"timeout"`

	// where to navigate
	Page string `mapstructure:"page" yaml:"page"`
}

// Load loads custom config from specified file or creates a new default one.
func Load(v *viper.Viper, file string) *Config {
	v.SetConfigFile(file)

	if err := v.ReadInConfig(); err != nil {
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			log.Fatalf("[Config - Load] Coulnd't read custom config file. %v.", err)
		}

		log.Printf(
			"[Config - Load] "+
				"- Couldn't load custom config file. %v. Creating default config...",
			err,
		)

		if err := os.MkdirAll(filepath.Dir(file), 0770); err != nil {
			log.Fatalf("[Config - Load] "+
				"- Couldn't create default config dir. %v", err)
		}

		f, err := os.Create(file)
		if err != nil {
			log.Fatalf("[Config - Load] "+
				"- Couldn't create default cofig file. %v", err)
		}

		defer f.Close()
		if _, err := f.Write(checkingautomata.DefaultConfig); err != nil {
			log.Fatalf("[Config - Load] "+
				"- Couldn't load default config file. %v", err)
		}
	}

	config := &Config{}
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("[Config - Load] "+
			"- Couldn't unmarshall config file. %v", err)
	}

	return config
}
