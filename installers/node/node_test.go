package node

import (
	"devbox/core"
	"devbox/internal/testutil"
	"devbox/system"
	"testing"
)

func TestNodeInstallRunsNVMAndNode(t *testing.T) {
	exec := system.NewFakeExecutor()

	ctx := &core.Context{
		Executor: exec,
		Logger:   &testutil.FakeLogger{},
	}

	tool := &NodeTool{}

	err := tool.Install(ctx, Options{
		NodeVersion: "20",
		SetDefault:  true,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(exec.Commands) == 0 {
		t.Fatal("expected commands to be executed")
	}
}
