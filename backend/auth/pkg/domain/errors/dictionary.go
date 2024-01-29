package errors

import (
	"fmt"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/errors"
)

type ErrCode int32

const (
	InvalidReqBind ErrCode = iota
	InvalidReqValidations
	ErrInvalidCredentials
	ErrRefreshToken
	defaultErr
)

var dictionary = map[ErrCode]*errors.CustomError{
	InvalidReqBind: {
		Code:    "invalid_request",
		Message: "the request content is not valid or has an invalid format",
	},
	InvalidReqValidations: {
		Code:    "invalid_request",
		Message: "the request does not contains all required fields",
	},
	ErrInvalidCredentials: {
		Code:    "invalid_credentials",
		Message: "the session cannot be generated because the credentials are invalids",
	},
	ErrRefreshToken: {
		Code:    "invalid_refresh_token",
		Message: "the session cannot be re generated because the refresh token is not valid",
	},
	defaultErr: {
		Code:    "generic_error",
		Message: "the server generates an unusual response, please contact to it support",
	},
}

func NewError(code ErrCode, details ...any) error {
	err, ok := dictionary[code]
	if !ok {
		err = dictionary[defaultErr]
	}

	for _, detail := range details {
		err.Details = append(err.Details, fmt.Sprintf("%v", detail))
	}

	return err
}
