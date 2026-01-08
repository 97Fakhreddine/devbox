package pm2

import (
	"devbox/core"
	"devbox/system"
	"testing"
)

func TestPM2IsInstalled(t *testing.T) {
	exec := system.NewFakeExecutor()
	exec.InstalledMap["pm2"] = true

	ctx := &core.Context{
		Executor: exec,
	}

	tool := &PM2Tool{}
	if !tool.IsInstalled(ctx) {
		t.Fatal("expected PM2 to be installed")
	}
}
