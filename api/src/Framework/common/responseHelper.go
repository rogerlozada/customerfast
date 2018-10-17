package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HTTPStatusCode status code
type HTTPStatusCode int

//
const (
	HTTPOk                 HTTPStatusCode = 200
	HTTPCreate             HTTPStatusCode = 201
	HTTPInavalid           HTTPStatusCode = 400
	HTTPNotAuthorized      HTTPStatusCode = 401
	HTTPFaltaPagamento     HTTPStatusCode = 402
	HTTPForbiden           HTTPStatusCode = 403
	HTTPNotFound           HTTPStatusCode = 404
	HTTPInternalError      HTTPStatusCode = 500
	HTTPNotImplemented     HTTPStatusCode = 501
	HTTPInvalidGatway      HTTPStatusCode = 502
	HTTPUnavailable        HTTPStatusCode = 503
	HTTPTimeout            HTTPStatusCode = 504
	HTTPNotSupportVerssion HTTPStatusCode = 505
)

//APIResult result response
type APIResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  string      `json:"errors"`
}

//ResponseError error response
func ResponseError(message string, codigoHTTP HTTPStatusCode, w http.ResponseWriter) http.ResponseWriter {
	apiResult := APIResult{Success: false, Errors: message}
	messageJSON, _ := json.Marshal(apiResult)

	codigo := int(codigoHTTP)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codigo)
	fmt.Fprintf(w, "%s", messageJSON)

	return w
}

//ResponseSuccess success response
func ResponseSuccess(any interface{}, HTTPCode HTTPStatusCode, w http.ResponseWriter) http.ResponseWriter {

	apiResult := APIResult{Success: true, Data: any}
	messageJSON, _ := json.Marshal(apiResult)
	code := int(HTTPCode)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", messageJSON)

	return w
}
