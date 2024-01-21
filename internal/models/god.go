package models

import (
	_ "github.com/lib/pq"
)

type God struct {
	Name     string   `json:"name"`
	Role     string   `json:"role"`
	WinRate  float32  `json:"win_rate"`
	PickRate float32  `json:"pick_rate"`
	BanRate  float32  `json:"ban_rate"`
	HotItems []string `json:"hot_items"`
	TopItems []string `json:"top_items"`
}
