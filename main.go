package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"redis-desktop-explorer/internal"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// Create application with options
	err = wails.Run(&options.App{
		Title:     app.appTitle,
		Width:     internal.MinWidth,
		Height:    internal.MinHeight,
		MinWidth:  internal.MinWidth,
		MinHeight: internal.MinHeight,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		OnBeforeClose:    app.beforeClose,
		Bind: []interface{}{
			app,
			app.connSrv,
			app.settingSrv,
		},
		Windows: app.windowsOptions,
		Mac:     app.macOptions,
		Linux:   app.linuxOptions,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
