// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Operator string

const (
	OperatorMtn    Operator = "mtn"
	OperatorOrange Operator = "orange"
)

func (e Operator) ToPointer() *Operator {
	return &e
}
func (e *Operator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mtn":
		fallthrough
	case "orange":
		*e = Operator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Operator: %v", v)
	}
}

type Type string

const (
	TypeCollection   Type = "collection"
	TypeDisbursement Type = "disbursement"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "collection":
		fallthrough
	case "disbursement":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Status string

const (
	StatusOperational Status = "OPERATIONAL"
	StatusSuspended   Status = "SUSPENDED"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OPERATIONAL":
		fallthrough
	case "SUSPENDED":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type Operation struct {
	Type   *Type   `json:"type,omitempty"`
	Status *Status `json:"status,omitempty"`
}

func (o *Operation) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Operation) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type Availability struct {
	Operator  *Operator  `json:"operator,omitempty"`
	Operation *Operation `json:"operation,omitempty"`
}

func (o *Availability) GetOperator() *Operator {
	if o == nil {
		return nil
	}
	return o.Operator
}

func (o *Availability) GetOperation() *Operation {
	if o == nil {
		return nil
	}
	return o.Operation
}
