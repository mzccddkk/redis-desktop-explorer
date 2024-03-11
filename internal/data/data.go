package data

import (
	"log"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewConnectionRepo,
	NewSettingRepo,
)

type Data struct {
}

func NewData() (*Data, func(), error) {
	cleanup := func() {
		log.Println("closing the data resources")
	}

	return &Data{}, cleanup, nil
}
