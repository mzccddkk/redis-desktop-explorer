package service

import (
	"context"

	"redis-desktop-explorer/internal/biz"
)

type SettingService struct {
	ctx context.Context
	uc  *biz.SettingUsecase
}

func NewSettingService(uc *biz.SettingUsecase) *SettingService {
	return &SettingService{
		uc: uc,
	}
}

func (s *SettingService) UpdateSetting(setting *biz.Setting) error {
	return s.uc.UpdateSetting(s.ctx, setting)
}
