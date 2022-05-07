// Code generated by go generate; DO NOT EDIT.
package images

import (
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *RemoveOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *RemoveOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAll set field All to given value
func (o *RemoveOptions) WithAll(value bool) *RemoveOptions {
	o.All = &value
	return o
}

// GetAll returns value of field All
func (o *RemoveOptions) GetAll() bool {
	if o.All == nil {
		var z bool
		return z
	}
	return *o.All
}

// WithForce set field Force to given value
func (o *RemoveOptions) WithForce(value bool) *RemoveOptions {
	o.Force = &value
	return o
}

// GetForce returns value of field Force
func (o *RemoveOptions) GetForce() bool {
	if o.Force == nil {
		var z bool
		return z
	}
	return *o.Force
}

// WithIgnore set field Ignore to given value
func (o *RemoveOptions) WithIgnore(value bool) *RemoveOptions {
	o.Ignore = &value
	return o
}

// GetIgnore returns value of field Ignore
func (o *RemoveOptions) GetIgnore() bool {
	if o.Ignore == nil {
		var z bool
		return z
	}
	return *o.Ignore
}
