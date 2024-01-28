package validator

import "github.com/go-playground/validator/v10"

type Validator interface {
	Validate(i interface{}) error
}

type validatorImpl struct{ v *validator.Validate }

func (v *validatorImpl) Validate(i interface{}) error {
	return v.v.Struct(i)
}

func NewValidator() Validator { return &validatorImpl{v: validator.New()} }
