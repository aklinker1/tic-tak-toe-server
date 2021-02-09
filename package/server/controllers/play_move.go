package controllers

import (
	"github.com/aklinker1/tic-tak-toe-server/package/ai"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/models"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi/operations"
	"github.com/aklinker1/tic-tak-toe-server/package/server/services"
	"github.com/go-openapi/runtime/middleware"
)

func UsePlayMoveController(api *operations.TicTakToeAPI) {
	api.PlayMoveHandler = operations.PlayMoveHandlerFunc(
		func(params operations.PlayMoveParams) middleware.Responder {
			game, responder := services.Move.MakeAMove(params.GameID, models.PlayerP1, *params.Position)
			if responder != nil {
				return responder
			}
			if game.Winner == nil {
				var err error
				game, err = ai.MakeAMove(game)
				if err != nil {
					return operations.NewPlayMoveDefault(500).WithPayload(err.Error())
				}
			}
			return operations.NewPlayMoveOK().WithPayload(services.Board.Sprint(game))
		},
	)
}
