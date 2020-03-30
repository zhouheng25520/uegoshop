package response

import (
	"uegoshop/errorcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response interface {
	// JSON serializes the given struct as JSON into the response body.
	// It also sets the Content-Type as "application/json".
	ReturnJsonError(message ...string)
	// JSON serializes the given struct as JSON into the response body.
	// It also sets the Content-Type as "application/json".
	ReturnJsonSuccess(data interface{})
	// XML serializes the given struct as XML into the response body.
	// It also sets the Content-Type as "application/xml".
	ReturnXmlError(message ...string)
	// XML serializes the given struct as XML into the response body.
	// It also sets the Content-Type as "application/xml".
	ReturnXmlSuccess(data interface{})
	// JSONP serializes the given struct as JSON into the response body.
	// It add padding to response body to request data from a server residing in a different domain than the client.
	// It also sets the Content-Type as "application/javascript".
	ReturnJsonpError(message ...string)
	// JSONP serializes the given struct as JSON into the response body.
	// It add padding to response body to request data from a server residing in a different domain than the client.
	// It also sets the Content-Type as "application/javascript".
	ReturnJsonpSuccess(data interface{})
}

type Body struct {
	Code    int         `json:"code" default:"200"`
	Message string      `json:"message" default:"success"`
	Data    interface{} `json:"data,omitempty" default:"" ""`
}

func NewResponseBody() *Body {
	return &Body{
		Code:    errorcode.OK.Code(),
		Message: errorcode.OK.Message(),
		Data:    []interface{}{},
	}
}

type Header struct {
	Context  *gin.Context
	HttpCode int
}

func NewResponseHeader(c *gin.Context) *Header {
	return &Header{
		Context:  c,
		HttpCode: http.StatusOK,
	}
}

type Response struct {
	Header *Header
	Body   *Body
}

//NewResponse
func NewResponse(c *gin.Context) *Response {
	return &Response{
		Header: NewResponseHeader(c),
		Body:   NewResponseBody(),
	}
}

//ReturnJsonError
func (r *Response) ReturnJsonError(message ...string) {

	r.setResponseBodyMessage(message...)

	r.ReturnJsonSuccess([]interface{}{})
}

//ReturnJsonSuccess
func (r *Response) ReturnJsonSuccess(data interface{}) {
	r.Body.Data = data
	r.Header.Context.JSON(r.Header.HttpCode, r.Body)
}

//ReturnXmlError
func (r *Response) ReturnXmlError(message ...string) {

	r.setResponseBodyMessage(message...)

	r.ReturnXmlSuccess([]interface{}{})
}

//ReturnXmlSuccess
func (r *Response) ReturnXmlSuccess(data interface{}) {
	r.Body.Data = data
	r.Header.Context.XML(r.Header.HttpCode, r.Body)
}

//ReturnJsonpError
func (r *Response) ReturnJsonpError(message ...string) {

	r.setResponseBodyMessage(message...)

	r.ReturnJsonpSuccess([]interface{}{})
}

//ReturnJsonpSuccess
func (r *Response) ReturnJsonpSuccess(data interface{}) {
	r.Body.Data = data
	r.Header.Context.JSONP(r.Header.HttpCode, r.Body)
}

// Set error message corresponding to error code
// if message params is not empty, well be used .
func (r *Response) setResponseBodyMessage(message ...string) {
	r.Body.Message = errorcode.Int(r.Body.Code).Message()
	if len(message) > 0 {
		r.Body.Message = message[0]
	}
}
