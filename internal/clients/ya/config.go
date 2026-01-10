package ya

type Config struct {
	Url      string
	Login    string
	Password string
	Token    string
}

func NewConfig(config Config) *Config {
	return &Config{
		"",
		"",
		"",
		"",
	}
}
