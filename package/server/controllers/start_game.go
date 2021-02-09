package controllers

import (
	"github.com/aklinker1/tic-tak-toe-server/package/server/board"
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi/operations"
	"github.com/aklinker1/tic-tak-toe-server/package/server/repos"
	"github.com/go-openapi/runtime/middleware"
)

func UseStartGameController(api *operations.TicTakToeAPI) {
	api.StartGameHandler = operations.StartGameHandlerFunc(
		func(params operations.StartGameParams) middleware.Responder {
			game, err := repos.Game.Create()
			if err != nil {
				return operations.NewStartGameDefault(500).WithPayload(err.Error())
			}
			return operations.NewStartGameCreated().WithPayload(board.Sprint(game))
		},
	)
}
