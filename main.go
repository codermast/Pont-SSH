package main

import (
	"PontSsh/backend/constant"
	"PontSsh/backend/service"
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	project := service.NewProject()
	connection := service.NewConnection()
	logInfo := service.NewLogInfo()
	config := service.NewConfig()

	app := &options.App{
		Title:     "Pont SSH 连接工具",
		Width:     1024,
		Height:    720,
		MinWidth:  constant.MIN_WINDOW_WIDTH,
		MinHeight: constant.MIN_WINDOW_HEIGHT,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			project.Startup(ctx)
			config.Startup(ctx)
			connection.Startup(ctx)
		},
		OnShutdown: func(ctx context.Context) {
			connection.Shutdown(ctx)
			config.Shutdown(ctx)
		},
		Bind: []interface{}{
			project,
			connection,
			logInfo,
			config,
		},
		// Mac 配置
		Mac: &mac.Options{
			TitleBar:   mac.TitleBarDefault(),
			Appearance: mac.DefaultAppearance,
			About: &mac.AboutInfo{
				Title:   "About",
				Message: "© 2024 CoderMast",
				Icon:    icon,
			},
		},
	}

	// Create application with options
	err := wails.Run(app)

	if err != nil {
		println("Error:", err.Error())
	}
}
