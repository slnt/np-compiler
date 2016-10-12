package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"

	"github.com/slantin/np-compiler/noonpacific"
	"github.com/slantin/np-compiler/soundcloud"
)

var config *Config

func init() {
	cfg, err := load()
	if err != nil {
		panic(err)
	}
	config = cfg
}

// Get returns the config that was loaded from config/config.yaml
func Get() *Config {
	return config
}

// Config is the aplication's configuration
type Config struct {
	SaveArtwork bool               `yaml:"save_artwork"`
	SoundCloud  soundcloud.Config  `yaml:"soundcloud"`
	NoonPacific noonpacific.Config `yaml:"noonpacific"`
}

func load() (*Config, error) {
	c := Config{}

	file, err := os.Open("./config/config.yaml")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	err = validator.Validate(c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
