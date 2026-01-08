package core

type Tool interface {
	ID() string
	Name() string
	Description() string

	IsInstalled(*Context) bool
	Install(*Context, InstallOptions) error
}
