// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	incident "github.com/mattermost/mattermost-plugin-incident-response/server/incident"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// CreateIncident provides a mock function with given fields: _a0
func (_m *Store) CreateIncident(_a0 *incident.Incident) (*incident.Incident, error) {
	ret := _m.Called(_a0)

	var r0 *incident.Incident
	if rf, ok := ret.Get(0).(func(*incident.Incident) *incident.Incident); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*incident.Incident)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*incident.Incident) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllHeaders provides a mock function with given fields:
func (_m *Store) GetAllHeaders() ([]incident.Header, error) {
	ret := _m.Called()

	var r0 []incident.Header
	if rf, ok := ret.Get(0).(func() []incident.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incident.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllIncidents provides a mock function with given fields:
func (_m *Store) GetAllIncidents() ([]incident.Incident, error) {
	ret := _m.Called()

	var r0 []incident.Incident
	if rf, ok := ret.Get(0).(func() []incident.Incident); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incident.Incident)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIncident provides a mock function with given fields: id
func (_m *Store) GetIncident(id string) (*incident.Incident, error) {
	ret := _m.Called(id)

	var r0 *incident.Incident
	if rf, ok := ret.Get(0).(func(string) *incident.Incident); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*incident.Incident)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NukeDB provides a mock function with given fields:
func (_m *Store) NukeDB() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateIncident provides a mock function with given fields: _a0
func (_m *Store) UpdateIncident(_a0 *incident.Incident) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*incident.Incident) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}