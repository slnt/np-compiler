package soundcloud

// Config is configuration data for the soundcloud API, retrieved from a yaml config file
type Config struct {
	// TODO: auth without password?
	Username string `yaml:"username" validate:"nonzero"`
	Password string `yaml:"password" validate:"nonzero"`
	Client   struct {
		ID     string `yaml:"id" validate:"nonzero"`
		Secret string `yaml:"secret" validate:"nonzero"`
	} `yaml:"client"`
}
