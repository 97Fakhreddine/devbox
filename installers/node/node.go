package node

import (
	"errors"
	"fmt"

	"devbox/core"
)

type NodeTool struct{}

func (n *NodeTool) ID() string          { return "node" }
func (n *NodeTool) Name() string        { return "Node.js (NVM)" }
func (n *NodeTool) Description() string { return "Install Node.js using NVM" }

func (n *NodeTool) IsInstalled(ctx *core.Context) bool {
	return ctx.Executor.IsInstalled("node")
}

func (n *NodeTool) Install(ctx *core.Context, opts core.InstallOptions) error {
	options, ok := opts.(Options)
	if !ok {
		return errors.New("invalid options for Node installer")
	}

	if options.NodeVersion == "" {
		return errors.New("node version is required")
	}

	if err := InstallNVM(ctx); err != nil {
		return err
	}

	cmd := fmt.Sprintf(`
		export NVM_DIR="$HOME/.nvm"
		[ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh"
		nvm install %s
	`, options.NodeVersion)

	if options.SetDefault {
		cmd += fmt.Sprintf("\n nvm alias default %s", options.NodeVersion)
	}

	if ctx.Logger != nil {
		ctx.Logger.Info("Installing Node " + options.NodeVersion)
	}

	return ctx.Executor.Run(cmd)
}
