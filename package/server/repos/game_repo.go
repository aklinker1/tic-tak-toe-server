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
		Status: "IN_PROGRESS",
	}
	err := repo.DB.Create(game).Error
	return game, err
}

func (repo *gameRepo) Read(id int64) (*entities.Game, error) {
	game := &entities.Game{}
	err := repo.DB.First(game, id).Error
	if err != nil {
		return nil, err
	}
	moves, err := Move.ReadForGame(game.ID)
	if err != nil {
		return nil, err
	}
	game.Moves = moves
	return game, nil
}

func (repo *gameRepo) Update(game *entities.Game) error {
	return repo.DB.Save(game).Error
}
