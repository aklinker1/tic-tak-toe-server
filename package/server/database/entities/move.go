package entities

import (
	"time"

	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
)

type Move struct {
	ID       int64          `json:"id" gorm:"primarykey"`
	Position uint           `json:"position"`
	PlayedBy *models.Player `json:"playedBy"`
	PlayedAt time.Time      `json:"playedAt"`
	GameID   uint64         `json:"omit"`
}
