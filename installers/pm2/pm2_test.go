package pm2

import (
	"devbox/core"
	"devbox/internal/testutil"
	"devbox/system"
	"testing"
)

func TestPM2InstallRunsNPM(t *testing.T) {
	exec := system.NewFakeExecutor()

	ctx := &core.Context{
		Executor: exec,
		Logger:   &testutil.FakeLogger{},
	}

	tool := &PM2Tool{}
	err := tool.Install(ctx, Options{})

	if err != nil {
		t.Fatal(err)
	}

	if exec.Commands[0] != "npm install -g pm2" {
		t.Fatalf("unexpected command: %s", exec.Commands[0])
	}
}
