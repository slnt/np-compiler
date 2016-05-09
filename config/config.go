package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"

	"github.com/slantin/np-compiler/soundcloud"
)

// Config is the aplication's configuration
type Config struct {
	SaveArtwork bool              `yaml:"saveArtwork"`
	SoundCloud  soundcloud.Config `yaml:"soundcloud"`
}

// Load loads config from config/config.yaml.
func Load() (*Config, error) {
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
