package entities

import (
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
)

type Game struct {
	ID     int64          `gorm:"primarykey"`
	Size   uint           `json:"size"`
	Status string         `json:"status"`
	Winner *models.Player `json:"winner"`
	Moves  []Move         `gorm:"foreignKey:GameID"`
}
