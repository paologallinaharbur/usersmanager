// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/paologallinaharbur/usersmanager/models"
)

// DeleteUserRequestProcessedCode is the HTTP code returned for type DeleteUserRequestProcessed
const DeleteUserRequestProcessedCode int = 203

/*DeleteUserRequestProcessed Deleted

swagger:response deleteUserRequestProcessed
*/
type DeleteUserRequestProcessed struct {
}

// NewDeleteUserRequestProcessed creates DeleteUserRequestProcessed with default headers values
func NewDeleteUserRequestProcessed() *DeleteUserRequestProcessed {

	return &DeleteUserRequestProcessed{}
}

// WriteResponse to the client
func (o *DeleteUserRequestProcessed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(203)
}

// DeleteUserBadRequestCode is the HTTP code returned for type DeleteUserBadRequest
const DeleteUserBadRequestCode int = 400

/*DeleteUserBadRequest bad request

swagger:response deleteUserBadRequest
*/
type DeleteUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserBadRequest creates DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {

	return &DeleteUserBadRequest{}
}

// WithPayload adds the payload to the delete user bad request response
func (o *DeleteUserBadRequest) WithPayload(payload *models.Error) *DeleteUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user bad request response
func (o *DeleteUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserInternalServerErrorCode is the HTTP code returned for type DeleteUserInternalServerError
const DeleteUserInternalServerErrorCode int = 500

/*DeleteUserInternalServerError internal server error

swagger:response deleteUserInternalServerError
*/
type DeleteUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserInternalServerError creates DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {

	return &DeleteUserInternalServerError{}
}

// WithPayload adds the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) WithPayload(payload *models.Error) *DeleteUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
