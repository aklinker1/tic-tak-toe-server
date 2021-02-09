package repos

import (
	"fmt"
	"time"

	"github.com/aklinker1/tic-tak-toe-server/package/server/database/entities"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
	"gorm.io/gorm"
)

type moveRepo struct {
	DB *gorm.DB
}

var Move = &moveRepo{}

func (repo *moveRepo) Setup(db *gorm.DB) {
	repo.DB = db
}

func (repo *moveRepo) CreateOnGame(game *entities.Game, position int64, player models.Player) error {
	move := entities.Move{
		Position: position,
		PlayedBy: &player,
		PlayedAt: time.Now(),
	}
	game.Moves = append(game.Moves, move)
	return Game.Update(game)
}

func (repo *moveRepo) ReadForGame(gameID int64) ([]entities.Move, error) {
	moves := []entities.Move{}
	err := repo.DB.Where("game_id = ?", gameID).Find(&moves).Error
	fmt.Println(moves)
	return moves, err
}
