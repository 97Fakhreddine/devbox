package core

type App struct {
	Ctx   *Context
	Tools []Tool
}

func NewApp(ctx *Context) *App {
	return &App{
		Ctx:   ctx,
		Tools: []Tool{},
	}
}

func (a *App) RegisterTool(tool Tool) {
	a.Tools = append(a.Tools, tool)
}

func (a *App) GetTools() []Tool {
	return a.Tools
}
