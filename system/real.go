package system

type RealExecutor struct{}

func (r *RealExecutor) Run(cmd string) error {
	return Run(cmd)
}

func (r *RealExecutor) IsInstalled(cmd string) bool {
	return IsInstalled(cmd)
}
