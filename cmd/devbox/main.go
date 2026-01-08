package main

import (
	"devbox/ui"
	"fmt"
	"os"

	"devbox/core"
	"devbox/installers/node"
	"devbox/installers/pm2"
	"devbox/system"

	"github.com/charmbracelet/bubbles/list"
)

// ---------- Simple Console Logger (Production) ----------

type ConsoleLogger struct{}

func (l *ConsoleLogger) Info(msg string) {
	fmt.Println("ℹ️ ", msg)
}

func (l *ConsoleLogger) Success(msg string) {
	fmt.Println("✅", msg)
}

func (l *ConsoleLogger) Error(msg string) {
	fmt.Println("❌", msg)
}

// -------------------------------------------------------

func main() {

	// ANSI color codes
	const (
		Reset   = "\033[0m"
		Red     = "\033[31m"
		Green   = "\033[32m"
		Yellow  = "\033[33m"
		Cyan    = "\033[36m"
		Blue    = "\033[34m"
		Magenta = "\033[35m"
	)
	// Big ASCII art banner
	fmt.Println(Cyan + `
		██████╗ ███████╗██╗   ██╗██████╗ ██╗  ██╗ ██████╗ ██████╗ 
		██╔══██╗██╔════╝██║   ██║██╔══██╗██║  ██║██╔═══██╗██╔══██╗
		██████╔╝█████╗  ██║   ██║██████╔╝███████║██║   ██║██████╔╝
		██╔═══╝ ██╔══╝  ██║   ██║██╔═══╝ ██╔══██║██║   ██║██╔═══╝ 
		██║     ███████╗╚██████╔╝██║     ██║  ██║╚██████╔╝██║     
		╚═╝     ╚══════╝ ╚═════╝ ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝     
		` + Reset)

	fmt.Println(Yellow + "⚡ Welcome to DevBox Installer ⚡" + Reset)
	fmt.Println(Magenta + "-----------------------------------" + Reset)

	fmt.Println(Green + "[1] Install Dev Tools" + Reset)
	fmt.Println(Green + "[2] Configure Environment" + Reset)
	fmt.Println(Green + "[3] Launch DevBox" + Reset)
	fmt.Println(Green + "[4] Exit" + Reset)
	fmt.Println(Magenta + "-----------------------------------" + Reset)
	// Safety: Linux only (for now)
	system.RequireLinux()

	// Global app context
	ctx := &core.Context{
		Config:   core.DefaultConfig(),
		Logger:   &ConsoleLogger{},
		Executor: &system.RealExecutor{},
	}

	// Register tools
	tools := []core.Tool{
		&node.NodeTool{},
		&pm2.PM2Tool{},
	}

	// ---------- Simple Menu ----------
	fmt.Println("⚡ DevBox Installer")
	fmt.Println("-------------------")

	// Convert to UI items
	var items []list.Item
	for _, t := range tools {
		items = append(items, ui.ToolItem{
			NameVal: t.Name(),
			IDVal:   t.ID(),
		})
	}
	model := ui.NewModel(items)
	selectedIdx, err := ui.Run(model)
	if err != nil {
		fmt.Println("Error running UI:", err)
		os.Exit(1)
	}

	if selectedIdx < 0 || selectedIdx >= len(tools) {
		fmt.Println("No tool selected")
		os.Exit(0)
	}

	selectedTool := tools[selectedIdx]

	// Installation
	switch tool := selectedTool.(type) {
	case *node.NodeTool:
		err := tool.Install(ctx, node.Options{
			NodeVersion: "20",
			SetDefault:  true,
		})
		if err != nil {
			ctx.Logger.Error(err.Error())
			os.Exit(1)
		}
		ctx.Logger.Success("Node.js installed successfully")

	case *pm2.PM2Tool:
		err := tool.Install(ctx, pm2.Options{})
		if err != nil {
			ctx.Logger.Error(err.Error())
			os.Exit(1)
		}
		ctx.Logger.Success("PM2 installed successfully")
	}
}
