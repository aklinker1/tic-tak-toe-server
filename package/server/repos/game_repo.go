package repos

import (
	"github.com/aklinker1/tic-tak-toe-server/package/server/database/entities"
	"gorm.io/gorm"
)

type gameRepo struct {
	DB *gorm.DB
}

var Game = &gameRepo{}

func (repo *gameRepo) Setup(db *gorm.DB) {
	repo.DB = db
}

func (repo *gameRepo) Create() (*entities.Game, error) {
	game := &entities.Game{
		Size:   3,
		Status: "IN_PROGRES",
	}
	err := repo.DB.Create(game).Error
	return game, err
}
