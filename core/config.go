package core

type Config struct {
	AppName      string
	Version      string
	DefaultShell string
	InstallDir   string
	Debug        bool
}

func DefaultConfig() *Config {
	return &Config{
		AppName:      "DevBox",
		Version:      "0.1.0",
		DefaultShell: "bash",
		InstallDir:   "$HOME/.devbox",
		Debug:        false,
	}
}
