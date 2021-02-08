// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PlayMoveOKCode is the HTTP code returned for type PlayMoveOK
const PlayMoveOKCode int = 200

/*PlayMoveOK The move was played or the game has ended

swagger:response playMoveOK
*/
type PlayMoveOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPlayMoveOK creates PlayMoveOK with default headers values
func NewPlayMoveOK() *PlayMoveOK {

	return &PlayMoveOK{}
}

// WithPayload adds the payload to the play move o k response
func (o *PlayMoveOK) WithPayload(payload string) *PlayMoveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the play move o k response
func (o *PlayMoveOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PlayMoveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PlayMoveBadRequestCode is the HTTP code returned for type PlayMoveBadRequest
const PlayMoveBadRequestCode int = 400

/*PlayMoveBadRequest An invalid move was played

swagger:response playMoveBadRequest
*/
type PlayMoveBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPlayMoveBadRequest creates PlayMoveBadRequest with default headers values
func NewPlayMoveBadRequest() *PlayMoveBadRequest {

	return &PlayMoveBadRequest{}
}

// WithPayload adds the payload to the play move bad request response
func (o *PlayMoveBadRequest) WithPayload(payload string) *PlayMoveBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the play move bad request response
func (o *PlayMoveBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PlayMoveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PlayMoveDefault Unknown Error

swagger:response playMoveDefault
*/
type PlayMoveDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPlayMoveDefault creates PlayMoveDefault with default headers values
func NewPlayMoveDefault(code int) *PlayMoveDefault {
	if code <= 0 {
		code = 500
	}

	return &PlayMoveDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the play move default response
func (o *PlayMoveDefault) WithStatusCode(code int) *PlayMoveDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the play move default response
func (o *PlayMoveDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the play move default response
func (o *PlayMoveDefault) WithPayload(payload string) *PlayMoveDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the play move default response
func (o *PlayMoveDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PlayMoveDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}