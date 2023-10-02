package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"tableflow/go/pkg/evaluator"
)

type ValidationType struct {
	Name      string              `json:"-" example:"regex"`
	Evaluator evaluator.Evaluator `json:"-"`
}

// Pre-defined ValidationTypes
var (
	ValidationNotBlank = ValidationType{"not_blank", evaluator.NotBlankEvaluator{}}
	ValidationRegex    = ValidationType{"regex", evaluator.RegexEvaluator{}}
)

func (v *ValidationType) Scan(value interface{}) error {
	// Convert the database value to a string
	typeStr, ok := value.(string)
	if !ok {
		return errors.New("failed to scan ValidationType")
	}
	// Set the Evaluator from the string type
	switch typeStr {
	case ValidationNotBlank.Name:
		*v = ValidationNotBlank
	case ValidationRegex.Name:
		*v = ValidationNotBlank
	default:
		return fmt.Errorf("The validation type %v is invalid", typeStr)
	}
	return nil
}

func (v ValidationType) Value() (driver.Value, error) {
	return v.Name, nil
}
