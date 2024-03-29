// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserBasic this user information is public
// swagger:model userBasic
type UserBasic struct {

	// id
	// Read Only: true
	// Format: ObjectId
	ID strfmt.ObjectId `json:"_id,omitempty"`

	// company
	// Max Length: 256
	Company string `json:"company,omitempty"`

	// description
	// Max Length: 4096
	Description string `json:"description,omitempty"`

	// full name
	// Required: true
	// Max Length: 256
	// Min Length: 10
	FullName *string `json:"fullName"`

	// gender
	// Required: true
	// Enum: [male female]
	Gender *string `json:"gender"`

	// images
	Images []*Images `json:"images"`

	// interest in
	// Required: true
	// Enum: [male female]
	InterestIn *string `json:"interestIn"`

	// job title
	// Max Length: 256
	JobTitle string `json:"jobTitle,omitempty"`

	// school
	// Max Length: 256
	School string `json:"school,omitempty"`
}

// Validate validates this user basic
func (m *UserBasic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompany(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterestIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserBasic) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("_id", "body", "ObjectId", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserBasic) validateCompany(formats strfmt.Registry) error {

	if swag.IsZero(m.Company) { // not required
		return nil
	}

	if err := validate.MaxLength("company", "body", string(m.Company), 256); err != nil {
		return err
	}

	return nil
}

func (m *UserBasic) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 4096); err != nil {
		return err
	}

	return nil
}

func (m *UserBasic) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("fullName", "body", m.FullName); err != nil {
		return err
	}

	if err := validate.MinLength("fullName", "body", string(*m.FullName), 10); err != nil {
		return err
	}

	if err := validate.MaxLength("fullName", "body", string(*m.FullName), 256); err != nil {
		return err
	}

	return nil
}

var userBasicTypeGenderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["male","female"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userBasicTypeGenderPropEnum = append(userBasicTypeGenderPropEnum, v)
	}
}

const (

	// UserBasicGenderMale captures enum value "male"
	UserBasicGenderMale string = "male"

	// UserBasicGenderFemale captures enum value "female"
	UserBasicGenderFemale string = "female"
)

// prop value enum
func (m *UserBasic) validateGenderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userBasicTypeGenderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserBasic) validateGender(formats strfmt.Registry) error {

	if err := validate.Required("gender", "body", m.Gender); err != nil {
		return err
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", *m.Gender); err != nil {
		return err
	}

	return nil
}

func (m *UserBasic) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var userBasicTypeInterestInPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["male","female"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userBasicTypeInterestInPropEnum = append(userBasicTypeInterestInPropEnum, v)
	}
}

const (

	// UserBasicInterestInMale captures enum value "male"
	UserBasicInterestInMale string = "male"

	// UserBasicInterestInFemale captures enum value "female"
	UserBasicInterestInFemale string = "female"
)

// prop value enum
func (m *UserBasic) validateInterestInEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userBasicTypeInterestInPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserBasic) validateInterestIn(formats strfmt.Registry) error {

	if err := validate.Required("interestIn", "body", m.InterestIn); err != nil {
		return err
	}

	// value enum
	if err := m.validateInterestInEnum("interestIn", "body", *m.InterestIn); err != nil {
		return err
	}

	return nil
}

func (m *UserBasic) validateJobTitle(formats strfmt.Registry) error {

	if swag.IsZero(m.JobTitle) { // not required
		return nil
	}

	if err := validate.MaxLength("jobTitle", "body", string(m.JobTitle), 256); err != nil {
		return err
	}

	return nil
}

func (m *UserBasic) validateSchool(formats strfmt.Registry) error {

	if swag.IsZero(m.School) { // not required
		return nil
	}

	if err := validate.MaxLength("school", "body", string(m.School), 256); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserBasic) UnmarshalBinary(b []byte) error {
	var res UserBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
