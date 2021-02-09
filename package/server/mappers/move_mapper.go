package mappers

import (
	"github.com/aklinker1/tic-tak-toe-server/package/server/database/entities"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
	"github.com/go-openapi/strfmt"
)

func MoveEntityArrayToModelArray(input []entities.Move) []*models.Move {
	output := []*models.Move{}
	for _, move := range input {
		output = append(output, MoveEntityToModel(move))
	}
	return output
}

func MoveEntityToModel(input entities.Move) *models.Move {
	return &models.Move{
		ID:       input.ID,
		PlayedAt: strfmt.DateTime(input.PlayedAt),
		PlayedBy: input.PlayedBy,
		Position: int64(input.Position),
	}
}
