package controllers

import (
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi/operations"
	"github.com/aklinker1/tic-tak-toe-server/package/server/repos"
	"github.com/go-openapi/runtime/middleware"
)

func UseQuitGameController(api *operations.TicTakToeAPI) {
	api.QuitGameHandler = operations.QuitGameHandlerFunc(
		func(params operations.QuitGameParams) middleware.Responder {
			game, err := repos.Game.Read(params.GameID)
			if err != nil {
				return operations.NewQuitGameDefault(500).WithPayload(err.Error())
			}
			if game.Status == "COMPLETED" {
				return operations.NewQuitGameBadRequest().WithPayload("The game is already over")
			}
			game.Status = "COMPLETED"
			err = repos.Game.Update(game)
			if err != nil {
				return operations.NewQuitGameDefault(500).WithPayload(err.Error())
			}
			return operations.NewQuitGameOK().WithPayload("Thanks for playing!")
		},
	)
}
