// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserInternal user internal
// swagger:model userInternal
type UserInternal struct {
	UserPrivate

	// deslikes
	// Read Only: true
	Deslikes []strfmt.ObjectId `json:"deslikes"`

	// facebook ID
	// Read Only: true
	FacebookID string `json:"facebookID,omitempty"`

	// likes
	// Read Only: true
	Likes []strfmt.ObjectId `json:"likes"`

	// store the subscription data
	// Read Only: true
	Subscription string `json:"subscription,omitempty"`

	// superlikes
	// Read Only: true
	Superlikes []strfmt.ObjectId `json:"superlikes"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UserInternal) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 UserPrivate
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.UserPrivate = aO0

	// AO1
	var dataAO1 struct {
		Deslikes []strfmt.ObjectId `json:"deslikes"`

		FacebookID string `json:"facebookID,omitempty"`

		Likes []strfmt.ObjectId `json:"likes"`

		Subscription string `json:"subscription,omitempty"`

		Superlikes []strfmt.ObjectId `json:"superlikes"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Deslikes = dataAO1.Deslikes

	m.FacebookID = dataAO1.FacebookID

	m.Likes = dataAO1.Likes

	m.Subscription = dataAO1.Subscription

	m.Superlikes = dataAO1.Superlikes

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UserInternal) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.UserPrivate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Deslikes []strfmt.ObjectId `json:"deslikes"`

		FacebookID string `json:"facebookID,omitempty"`

		Likes []strfmt.ObjectId `json:"likes"`

		Subscription string `json:"subscription,omitempty"`

		Superlikes []strfmt.ObjectId `json:"superlikes"`
	}

	dataAO1.Deslikes = m.Deslikes

	dataAO1.FacebookID = m.FacebookID

	dataAO1.Likes = m.Likes

	dataAO1.Subscription = m.Subscription

	dataAO1.Superlikes = m.Superlikes

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this user internal
func (m *UserInternal) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UserPrivate
	if err := m.UserPrivate.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeslikes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLikes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuperlikes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserInternal) validateDeslikes(formats strfmt.Registry) error {

	if swag.IsZero(m.Deslikes) { // not required
		return nil
	}

	for i := 0; i < len(m.Deslikes); i++ {

		if err := validate.FormatOf("deslikes"+"."+strconv.Itoa(i), "body", "ObjectId", m.Deslikes[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *UserInternal) validateLikes(formats strfmt.Registry) error {

	if swag.IsZero(m.Likes) { // not required
		return nil
	}

	for i := 0; i < len(m.Likes); i++ {

		if err := validate.FormatOf("likes"+"."+strconv.Itoa(i), "body", "ObjectId", m.Likes[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *UserInternal) validateSuperlikes(formats strfmt.Registry) error {

	if swag.IsZero(m.Superlikes) { // not required
		return nil
	}

	for i := 0; i < len(m.Superlikes); i++ {

		if err := validate.FormatOf("superlikes"+"."+strconv.Itoa(i), "body", "ObjectId", m.Superlikes[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserInternal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserInternal) UnmarshalBinary(b []byte) error {
	var res UserInternal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
