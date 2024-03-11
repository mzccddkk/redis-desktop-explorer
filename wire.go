//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"redis-desktop-explorer/internal/biz"
	"redis-desktop-explorer/internal/data"
	"redis-desktop-explorer/internal/service"
)

// wireApp init application.
func wireApp() (*App, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		NewApp,
	))
}
