package config

// Config is the aplication's configuration
type Config struct {
	SoundCloud SoundCloudConfig `yaml:"soundcloud"`
}

// SoundCloudConfig defines the configuration for SoudnCloud requests
type SoundCloudConfig struct {
	ClientID     string `yaml:"ID"`
	ClientSecret string `yaml:"secret"`
}

// Load loads config from config/config.yaml.
func Load() *Config {
	c := Config{}

	return &c
}
