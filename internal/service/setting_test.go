package service

import (
	"testing"

	"redis-desktop-explorer/internal/biz"
	"redis-desktop-explorer/internal/data"
)

func newSettingService() *SettingService {
	dataData, _, _ := data.NewData()
	settingRepo := data.NewSettingRepo(dataData)
	settingUsecase := biz.NewSettingUsecase(settingRepo)
	return NewSettingService(settingUsecase)
}

func TestUpdateSetting(t *testing.T) {
	settingService := newSettingService()

	setting := &biz.Setting{
		Width:  1024,
		Height: 633,
	}
	err := settingService.UpdateSetting(setting)
	if err != nil {
		t.Errorf("update failed: %v", err)
	}
}
