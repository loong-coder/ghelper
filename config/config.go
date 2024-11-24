package config

type Config struct {
	Author  string `yaml:"author"`
	License string `yaml:"license"`
	Overite []struct {
		Source string `yaml:"source"`
		Target string `yaml:"target"`
	}
}
