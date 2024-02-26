package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"redis-desktop-explorer/internal"
)

// App struct
type App struct {
	ctx            context.Context
	appName        string
	windowsOptions *windows.Options
	macOptions     *mac.Options
	linuxOptions   *linux.Options
}

// NewApp creates a new App application struct
func NewApp() *App {
	a := &App{}
	a.appName = "Redis Desktop Explorer"
	a.windowsOptions = nil
	a.macOptions = &mac.Options{
		About: &mac.AboutInfo{
			Title:   a.appName,
			Message: "Version: " + internal.Version,
		},
	}
	a.linuxOptions = nil
	return a
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	// todo
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
