package biz

import (
	"context"
)

type Setting struct {
	Width  int `json:"width,omitempty" yaml:"width,omitempty"`
	Height int `json:"height,omitempty" yaml:"height,omitempty"`
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
