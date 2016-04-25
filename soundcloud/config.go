package soundcloud

type Config struct {
	// TODO: can I auth without storing password in plaintext? ;_;
	Password     string `yaml:"password" validate:"nonzero"`
	ClientID     string `yaml:"id" validate:"nonzero"`
	ClientSecret string `yaml:"secret" validate:"nonzero"`
}
