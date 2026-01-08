package system

type Executor interface {
	Run(cmd string) error
	IsInstalled(cmd string) bool
}
