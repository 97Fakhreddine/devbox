package system

import "fmt"

type FakeExecutor struct {
	Commands     []string
	InstalledMap map[string]bool
	FailOn       string
}

func NewFakeExecutor() *FakeExecutor {
	return &FakeExecutor{
		InstalledMap: map[string]bool{},
	}
}

func (f *FakeExecutor) Run(cmd string) error {
	f.Commands = append(f.Commands, cmd)
	if f.FailOn != "" && f.FailOn == cmd {
		return fmt.Errorf("forced failure")
	}
	return nil
}

func (f *FakeExecutor) IsInstalled(cmd string) bool {
	return f.InstalledMap[cmd]
}
