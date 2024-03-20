package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"redis-desktop-explorer/internal"
	"redis-desktop-explorer/internal/service"
)

// App struct
type App struct {
	ctx            context.Context
	appName        string
	appTitle       string
	windowsOptions *windows.Options
	macOptions     *mac.Options
	linuxOptions   *linux.Options

	// services
	connSrv    *service.ConnectionService
	settingSrv *service.SettingService
}

// NewApp creates a new App application struct
func NewApp(connSrv *service.ConnectionService, settingSrv *service.SettingService) *App {
	return &App{
		ctx:            nil,
		appName:        internal.AppName,
		appTitle:       internal.AppName + " " + internal.Version,
		windowsOptions: nil,
		macOptions: &mac.Options{
			About: &mac.AboutInfo{
				Title:   internal.AppName,
				Message: "Version: " + internal.Version,
			},
		},
		linuxOptions: nil,
		connSrv:      connSrv,
		settingSrv:   settingSrv,
	}
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	// todo
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	w, h := runtime.WindowGetSize(ctx)
	log.Println("Window size: ", w, h)
	return false
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
