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

	app := &options.App{
		Title:             "Pont SSH 连接工具",
		Width:             1024,
		Height:            720,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		MinWidth:          constant.MIN_WINDOW_WIDTH,
		MinHeight:         constant.MIN_WINDOW_HEIGHT,
		MaxWidth:          0,
		MaxHeight:         0,
		StartHidden:       false,
		HideWindowOnClose: false,
		AlwaysOnTop:       false,
		BackgroundColour:  nil,
		Assets:            nil,
		AssetsHandler:     nil,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:               nil,
		Logger:             nil,
		LogLevel:           0,
		LogLevelProduction: 0,
		OnStartup: func(ctx context.Context) {
			project.Startup(ctx)
			connection.Startup(ctx)
		},
		OnDomReady:    nil,
		OnShutdown:    connection.Shutdown,
		OnBeforeClose: nil,
		Bind: []interface{}{
			project,
			connection,
			logInfo,
		},
		EnumBind:                         nil,
		WindowStartState:                 0,
		ErrorFormatter:                   nil,
		CSSDragProperty:                  "",
		CSSDragValue:                     "",
		EnableDefaultContextMenu:         false,
		EnableFraudulentWebsiteDetection: false,
		SingleInstanceLock:               nil,
		Windows:                          nil,
		// Mac 配置
		Mac: &mac.Options{
			TitleBar:   mac.TitleBarDefault(),
			Appearance: mac.DefaultAppearance,
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
		Linux:        nil,
		Experimental: nil,
		Debug:        options.Debug{},
		DragAndDrop:  nil,
	}

	// Create application with options
	err := wails.Run(app)

	if err != nil {
		println("Error:", err.Error())
	}
}
