// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewPlayMoveParams creates a new PlayMoveParams object
// no default values defined in spec.
func NewPlayMoveParams() PlayMoveParams {

	return PlayMoveParams{}
}

// PlayMoveParams contains all the bound params for the play move operation
// typically these are obtained from a http.Request
//
// swagger:parameters playMove
type PlayMoveParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The id of the game that will be effected by the operation
	  Required: true
	  In: path
	*/
	GameID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPlayMoveParams() beforehand.
func (o *PlayMoveParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGameID, rhkGameID, _ := route.Params.GetOK("gameId")
	if err := o.bindGameID(rGameID, rhkGameID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGameID binds and validates parameter GameID from path.
func (o *PlayMoveParams) bindGameID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GameID = raw

	return nil
}
