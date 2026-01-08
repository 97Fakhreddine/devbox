package node

import "devbox/core"

func InstallNVM(ctx *core.Context) error {
	if ctx.Executor.IsInstalled("nvm") {
		return nil
	}

	if ctx.Logger != nil {
		ctx.Logger.Info("Installing NVM...")
	}

	return ctx.Executor.Run(`
		curl -fsSL https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
	`)
}
