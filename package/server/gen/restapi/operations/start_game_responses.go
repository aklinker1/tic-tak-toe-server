// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// StartGameCreatedCode is the HTTP code returned for type StartGameCreated
const StartGameCreatedCode int = 201

/*StartGameCreated Game was started successfully

swagger:response startGameCreated
*/
type StartGameCreated struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewStartGameCreated creates StartGameCreated with default headers values
func NewStartGameCreated() *StartGameCreated {

	return &StartGameCreated{}
}

// WithPayload adds the payload to the start game created response
func (o *StartGameCreated) WithPayload(payload string) *StartGameCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the start game created response
func (o *StartGameCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StartGameCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*StartGameDefault Unknown Error

swagger:response startGameDefault
*/
type StartGameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewStartGameDefault creates StartGameDefault with default headers values
func NewStartGameDefault(code int) *StartGameDefault {
	if code <= 0 {
		code = 500
	}

	return &StartGameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the start game default response
func (o *StartGameDefault) WithStatusCode(code int) *StartGameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the start game default response
func (o *StartGameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the start game default response
func (o *StartGameDefault) WithPayload(payload string) *StartGameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the start game default response
func (o *StartGameDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StartGameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
