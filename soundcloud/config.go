package soundcloud

type Config struct {
	// TODO: can I auth without storing password in plaintext? ;_;
	Password     string `yaml:"password"`
	ClientID     string `yaml:"id"`
	ClientSecret string `yaml:"secret"`
}
