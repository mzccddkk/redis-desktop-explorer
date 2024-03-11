package biz

import (
	"context"
)

type Setting struct {
	Width  int `json:"width" yaml:"width"`
	Height int `json:"height" yaml:"height"`
}

type SettingRepo interface {
	Update(ctx context.Context, setting *Setting) error
}

type SettingUsecase struct {
	repo SettingRepo
}

func NewSettingUsecase(repo SettingRepo) *SettingUsecase {
	return &SettingUsecase{
		repo: repo,
	}
}

func (uc *SettingUsecase) UpdateSetting(ctx context.Context, setting *Setting) error {
	return uc.repo.Update(ctx, setting)
}
