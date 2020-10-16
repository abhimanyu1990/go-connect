// Package classification Forum.
//
// Documentation of Forum API.
//
//     Schemes: http
//     BasePath: /api
//     Version: 1.0.0
//     Host: 192.168.0.119:8888
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
//     SecurityDefinitions:
//       api_key:
//            type: apiKey
//            name: Authorization
//            in: header
//
// swagger:meta
package docs

import "go-connect/app/dto"
// swagger:route POST /register register idOfRegisterEndpoint
// User registration.
// responses: 
//   200: UserResponse
//   401: ResponseError
//   400: ResponseError
//   default: ResponseError
// This text will appear as description of your response body.
// swagger:response UserResponse
type regsiterResponseWrapper struct {
	// in:body
	Body dto.UserResponse
}

// This text will appear as description of your response body.
// swagger:response ResponseError
type regsiterErrorResponseWrapper struct {
	// This text will appear as description of your error body.
	// in:body
	Body dto.ResponseError
}

// swagger:parameters idOfRegisterEndpoint
type regsiterRequestWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body dto.UserRequest
}