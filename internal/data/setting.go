package data

import (
	"context"

	"gopkg.in/yaml.v2"
	"redis-desktop-explorer/internal"
	"redis-desktop-explorer/internal/biz"
)

type settingRepo struct {
	data         *Data
	localStorage *internal.Storage
}

func NewSettingRepo(data *Data) biz.SettingRepo {
	return &settingRepo{
		data:         data,
		localStorage: internal.NewStorage(internal.SettingsFile),
	}
}

func (r *settingRepo) Update(ctx context.Context, setting *biz.Setting) error {
	data, err := yaml.Marshal(setting)
	if err != nil {
		return err
	}

	return r.localStorage.Put(data)
}
