// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// StartGameHandlerFunc turns a function with the right signature into a start game handler
type StartGameHandlerFunc func(StartGameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn StartGameHandlerFunc) Handle(params StartGameParams) middleware.Responder {
	return fn(params)
}

// StartGameHandler interface for that can handle valid start game params
type StartGameHandler interface {
	Handle(StartGameParams) middleware.Responder
}

// NewStartGame creates a new http.Handler for the start game operation
func NewStartGame(ctx *middleware.Context, handler StartGameHandler) *StartGame {
	return &StartGame{Context: ctx, Handler: handler}
}

/*StartGame swagger:route POST /games startGame

Start a new game

*/
type StartGame struct {
	Context *middleware.Context
	Handler StartGameHandler
}

func (o *StartGame) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewStartGameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
