// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/paologallinaharbur/usersmanager/models"
)

// UpdateUserAcceptedCode is the HTTP code returned for type UpdateUserAccepted
const UpdateUserAcceptedCode int = 202

/*UpdateUserAccepted OK

swagger:response updateUserAccepted
*/
type UpdateUserAccepted struct {

	/*
	  In: Body
	*/
	Payload models.NickName `json:"body,omitempty"`
}

// NewUpdateUserAccepted creates UpdateUserAccepted with default headers values
func NewUpdateUserAccepted() *UpdateUserAccepted {

	return &UpdateUserAccepted{}
}

// WithPayload adds the payload to the update user accepted response
func (o *UpdateUserAccepted) WithPayload(payload models.NickName) *UpdateUserAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user accepted response
func (o *UpdateUserAccepted) SetPayload(payload models.NickName) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateUserBadRequestCode is the HTTP code returned for type UpdateUserBadRequest
const UpdateUserBadRequestCode int = 400

/*UpdateUserBadRequest bad request

swagger:response updateUserBadRequest
*/
type UpdateUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserBadRequest creates UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {

	return &UpdateUserBadRequest{}
}

// WithPayload adds the payload to the update user bad request response
func (o *UpdateUserBadRequest) WithPayload(payload *models.Error) *UpdateUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user bad request response
func (o *UpdateUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserInternalServerErrorCode is the HTTP code returned for type UpdateUserInternalServerError
const UpdateUserInternalServerErrorCode int = 500

/*UpdateUserInternalServerError internal server error

swagger:response updateUserInternalServerError
*/
type UpdateUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserInternalServerError creates UpdateUserInternalServerError with default headers values
func NewUpdateUserInternalServerError() *UpdateUserInternalServerError {

	return &UpdateUserInternalServerError{}
}

// WithPayload adds the payload to the update user internal server error response
func (o *UpdateUserInternalServerError) WithPayload(payload *models.Error) *UpdateUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user internal server error response
func (o *UpdateUserInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
