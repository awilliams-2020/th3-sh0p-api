// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostImageOKCode is the HTTP code returned for type PostImageOK
const PostImageOKCode int = 200

/*PostImageOK Images to show

swagger:response postImageOK
*/
type PostImageOK struct {

	/*
	  In: Body
	*/
	Payload *PostImageOKBody `json:"body,omitempty"`
}

// NewPostImageOK creates PostImageOK with default headers values
func NewPostImageOK() *PostImageOK {

	return &PostImageOK{}
}

// WithPayload adds the payload to the post image o k response
func (o *PostImageOK) WithPayload(payload *PostImageOKBody) *PostImageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post image o k response
func (o *PostImageOK) SetPayload(payload *PostImageOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostImageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostImageBadRequestCode is the HTTP code returned for type PostImageBadRequest
const PostImageBadRequestCode int = 400

/*PostImageBadRequest openAI returned a non 200

swagger:response postImageBadRequest
*/
type PostImageBadRequest struct {
}

// NewPostImageBadRequest creates PostImageBadRequest with default headers values
func NewPostImageBadRequest() *PostImageBadRequest {

	return &PostImageBadRequest{}
}

// WriteResponse to the client
func (o *PostImageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostImageInternalServerErrorCode is the HTTP code returned for type PostImageInternalServerError
const PostImageInternalServerErrorCode int = 500

/*PostImageInternalServerError couldn't process request

swagger:response postImageInternalServerError
*/
type PostImageInternalServerError struct {
}

// NewPostImageInternalServerError creates PostImageInternalServerError with default headers values
func NewPostImageInternalServerError() *PostImageInternalServerError {

	return &PostImageInternalServerError{}
}

// WriteResponse to the client
func (o *PostImageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
