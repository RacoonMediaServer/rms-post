// Code generated by go-swagger; DO NOT EDIT.

package notify

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// NotifyEmailOKCode is the HTTP code returned for type NotifyEmailOK
const NotifyEmailOKCode int = 200

/*
NotifyEmailOK OK

swagger:response notifyEmailOK
*/
type NotifyEmailOK struct {
}

// NewNotifyEmailOK creates NotifyEmailOK with default headers values
func NewNotifyEmailOK() *NotifyEmailOK {

	return &NotifyEmailOK{}
}

// WriteResponse to the client
func (o *NotifyEmailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// NotifyEmailBadRequestCode is the HTTP code returned for type NotifyEmailBadRequest
const NotifyEmailBadRequestCode int = 400

/*
NotifyEmailBadRequest Ошибка во входных параметрах

swagger:response notifyEmailBadRequest
*/
type NotifyEmailBadRequest struct {
}

// NewNotifyEmailBadRequest creates NotifyEmailBadRequest with default headers values
func NewNotifyEmailBadRequest() *NotifyEmailBadRequest {

	return &NotifyEmailBadRequest{}
}

// WriteResponse to the client
func (o *NotifyEmailBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// NotifyEmailInternalServerErrorCode is the HTTP code returned for type NotifyEmailInternalServerError
const NotifyEmailInternalServerErrorCode int = 500

/*
NotifyEmailInternalServerError Ошибка на стороне сервера

swagger:response notifyEmailInternalServerError
*/
type NotifyEmailInternalServerError struct {
}

// NewNotifyEmailInternalServerError creates NotifyEmailInternalServerError with default headers values
func NewNotifyEmailInternalServerError() *NotifyEmailInternalServerError {

	return &NotifyEmailInternalServerError{}
}

// WriteResponse to the client
func (o *NotifyEmailInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
