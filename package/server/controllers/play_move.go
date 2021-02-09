package controllers

import (
	"github.com/aklinker1/tic-tak-toe-server/package/server/board"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi/operations"
	"github.com/aklinker1/tic-tak-toe-server/package/server/repos"
	"github.com/go-openapi/runtime/middleware"
)

func UsePlayMoveController(api *operations.TicTakToeAPI) {
	api.PlayMoveHandler = operations.PlayMoveHandlerFunc(
		func(params operations.PlayMoveParams) middleware.Responder {
			game, err := repos.Game.Read(params.GameID)
			if err != nil {
				return operations.NewPlayMoveNotFound()
			}
			if game.Status == "COMPLETED" {
				return operations.NewPlayMoveBadRequest().WithPayload("Game is already finished")
			}
			err = repos.Move.CreateOnGame(game, *params.Position, models.PlayerP1)
			if err != nil {
				return operations.NewPlayMoveDefault(500).WithPayload(err.Error())
			}
			newGame, err := repos.Game.Read(params.GameID)
			if err != nil {
				return operations.NewPlayMoveNotFound()
			}
			return operations.NewPlayMoveOK().WithPayload(board.Sprint(newGame))
		},
	)
}
