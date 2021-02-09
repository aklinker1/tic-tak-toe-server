package services

import (
	"fmt"

	"github.com/aklinker1/tic-tak-toe-server/package/server/database/entities"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi/operations"
	"github.com/aklinker1/tic-tak-toe-server/package/server/repos"
	"github.com/go-openapi/runtime/middleware"
)

type moveService struct {
}

var Move = &moveService{}

func (service *moveService) MakeAMove(gameID int64, player models.Player, position int64) (game *entities.Game, responder middleware.Responder) {
	game, err := repos.Game.Read(gameID)
	if err != nil {
		return nil, operations.NewPlayMoveNotFound()
	}
	if game.Status == "COMPLETED" {
		return nil, operations.NewPlayMoveBadRequest().WithPayload(fmt.Sprintf("Game %d (%d) is already finished", game.ID, gameID))
	}
	if !Board.AvailablePositions(game)[position] {
		return nil, operations.NewPlayMoveBadRequest().WithPayload("That position is already taken")
	}
	err = repos.Move.CreateOnGame(game, position, player)
	if err != nil {
		return nil, operations.NewPlayMoveDefault(500).WithPayload(err.Error())
	}
	newGame, err := repos.Game.Read(gameID)
	if err != nil {
		return nil, operations.NewPlayMoveNotFound()
	}

	winner := Board.FindWinner(newGame)
	if winner != nil {
		newGame.Winner = winner
		newGame.Status = "COMPLETED"
		err = repos.Game.Update(newGame)
		if err != nil {
			return nil, operations.NewPlayMoveDefault(500).WithPayload(err.Error())
		}
	}
	return newGame, nil
}
