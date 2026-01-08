package installers

import "devbox/core"

type BaseInstaller struct{}

func (b *BaseInstaller) RequireLinux() error {
	// placeholder for OS validation
	return nil
}

func (b *BaseInstaller) LogStart(ctx *core.Context, name string) {
	ctx.Logger.Info("Installing " + name + "...")
}
