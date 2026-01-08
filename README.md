# devbox

The project Architecture
```
devbox/
├── cmd/
│   └── devbox/
│       └── main.go          # App entry point
│
├── core/                    # Central brain (config, contracts)
│   ├── app.go
│   ├── config.go
│   ├── context.go
│   ├── tool.go
│   └── logger.go
│
├── ui/                      # TUI layer (Bubble Tea)
│   ├── model.go
│   ├── menu.go
│   ├── styles.go
│   └── messages.go
│
├── installers/              # Tool installers (plugins)
│   ├── installer.go
│   ├── node/
│   │   ├── node.go
│   │   └── nvm.go
│   └── pm2/
│       └── pm2.go
```
