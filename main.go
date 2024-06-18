package main

import (
	"PontSsh/backend/constant"
	"PontSsh/backend/service"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed all:backend/database/pontssh.db
var database embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {

	connection := service.NewConnection()

	app := &options.App{
		Title:     "Pont SSH 连接工具",
		Width:     1024,
		Height:    720,
		MinWidth:  constant.MIN_WINDOW_WIDTH,
		MinHeight: constant.MIN_WINDOW_HEIGHT,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  connection.Startup,
		OnShutdown: connection.Shutdown,
		Bind: []interface{}{
			connection,
		},

		// Mac 配置
		Mac: &mac.Options{
			TitleBar:   mac.TitleBarHiddenInset(),
			Appearance: mac.NSAppearanceNameDarkAqua,
			About: &mac.AboutInfo{
				Title:   "About",
				Message: "© 2021 Me",
				Icon:    icon,
			},

			Preferences: &mac.Preferences{
				TabFocusesLinks:        mac.Enabled,
				TextInteractionEnabled: mac.Disabled,
				FullscreenEnabled:      mac.Disabled,
			},
		},
	}

	// Create application with options
	err := wails.Run(app)

	if err != nil {
		println("Error:", err.Error())
	}
}
