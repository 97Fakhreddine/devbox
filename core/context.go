package core

import "devbox/system"

type Context struct {
	Config   *Config
	Logger   Logger
	Executor system.Executor
}
