package controllers

import (
	"github.com/aklinker1/tic-tak-toe-server/package/server/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func UseHealthController(api *operations.TicTakToeAPI) {
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(
		func(params operations.CheckHealthParams) middleware.Responder {
			return operations.NewCheckHealthNoContent()
		},
	)
}
