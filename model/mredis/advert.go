package mredis

import (
	"time"
	"uegoshop/model"
)

type Advert struct {
	ID           int                 `json:"id"`
	Name         string              `json:"name"`
	Code         string              `json:"code"`
	EnabledCycle int                 `json:"enabled_cycle"`
	IsEnabled    int                 `json:"is_enabled"`
	FromDate     time.Time           `json:"-"`
	FromDateTime int64               `json:"from_date_time"`
	ExpireTime   int64               `json:"expired"`
	Items        []*model.AdvertItem `json:"items"`
	ToDate       time.Time           `json:"-"`
	UpdatedAt    time.Time           `json:"-"`
	CreatedAt    time.Time           `json:"-"`
}
