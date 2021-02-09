package entities

import (
	"time"

	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
)

type Move struct {
	ID       int64          `json:"id" gorm:"primarykey"`
	Position int64          `json:"position"`
	PlayedBy *models.Player `json:"playedBy"`
	PlayedAt time.Time      `json:"playedAt"`
	GameID   int64          `json:"omit"`
}
