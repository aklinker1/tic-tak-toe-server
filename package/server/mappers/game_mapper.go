package mappers

import (
	"github.com/aklinker1/tic-tak-toe-server/package/server/database/entities"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
)

func GameEntityToModel(input *entities.Game) *models.Game {
	return &models.Game{
		ID:     input.ID,
		Moves:  MoveEntityArrayToModelArray(input.Moves),
		Size:   int64(input.Size),
		Status: input.Status,
		Winner: input.Winner,
	}
}
