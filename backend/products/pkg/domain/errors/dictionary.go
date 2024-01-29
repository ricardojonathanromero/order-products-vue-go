package errors

import (
	"fmt"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/errors"
)

type ErrCode int32

const (
	InvalidReqBind ErrCode = iota
	InvalidReqValidations
	InvalidId
	ErrNoDocument
	ErrNoDocuments
	ErrBindDBDocuments
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
	InvalidId: {
		Code:    "invalid_request",
		Message: "the field 'id' is required and must be type 'primitive.ObjectId'",
	},
	ErrNoDocument: {
		Code:    "document_not_found",
		Message: "the document id is not valid or not exists",
	},
	ErrNoDocuments: {
		Code:    "documents_not_found",
		Message: "the filters applied does not return any result",
	},
	ErrBindDBDocuments: {
		Code:    "db_error_struct",
		Message: "an error occurs trying to parse document in struct",
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
