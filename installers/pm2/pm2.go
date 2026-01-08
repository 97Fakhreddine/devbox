package pm2

import "devbox/core"

type PM2Tool struct{}

func (p *PM2Tool) ID() string          { return "pm2" }
func (p *PM2Tool) Name() string        { return "PM2" }
func (p *PM2Tool) Description() string { return "Node.js process manager" }

func (p *PM2Tool) IsInstalled(ctx *core.Context) bool {
	return ctx.Executor.IsInstalled("pm2")
}

func (p *PM2Tool) Install(ctx *core.Context, _ core.InstallOptions) error {
	if ctx.Logger != nil {
		ctx.Logger.Info("Installing PM2...")
	}
	return ctx.Executor.Run("npm install -g pm2")
}
