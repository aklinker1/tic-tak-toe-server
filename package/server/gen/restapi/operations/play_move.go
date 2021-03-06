// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PlayMoveHandlerFunc turns a function with the right signature into a play move handler
type PlayMoveHandlerFunc func(PlayMoveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PlayMoveHandlerFunc) Handle(params PlayMoveParams) middleware.Responder {
	return fn(params)
}

// PlayMoveHandler interface for that can handle valid play move params
type PlayMoveHandler interface {
	Handle(PlayMoveParams) middleware.Responder
}

// NewPlayMove creates a new http.Handler for the play move operation
func NewPlayMove(ctx *middleware.Context, handler PlayMoveHandler) *PlayMove {
	return &PlayMove{Context: ctx, Handler: handler}
}

/*PlayMove swagger:route POST /games/{gameId}/moves playMove

Play a move

*/
type PlayMove struct {
	Context *middleware.Context
	Handler PlayMoveHandler
}

func (o *PlayMove) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPlayMoveParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
