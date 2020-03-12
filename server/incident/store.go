package incident

import (
	pluginapi "github.com/mattermost/mattermost-plugin-api"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

const (
	allHeadersKey = "all_headers"
	incidentKey   = "incident_"
)

type idHeaderMap map[string]Header

var _ Store = &StoreImpl{}

// StoreImpl Implements incident store interface.
type StoreImpl struct {
	pluginAPI *pluginapi.Client
}

// NewStore creates a new store for incident service.
func NewStore(pluginAPI *pluginapi.Client) *StoreImpl {
	newStore := &StoreImpl{
		pluginAPI: pluginAPI,
	}
	return newStore
}

// GetAllHeaders Creates a new incident.
func (s *StoreImpl) GetAllHeaders() ([]Header, error) {
	headers, err := s.getIDHeaders()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all headers value")
	}

	return toHeader(headers), nil
}

// CreateIncident Creates a new incident.
func (s *StoreImpl) CreateIncident(incident *Incident) (*Incident, error) {
	if incident == nil {
		return nil, errors.New("incident is nil")
	}
	if incident.ID != "" {
		return nil, errors.New("ID should not be set")
	}
	incident.ID = model.NewId()

	saved, err := s.pluginAPI.KV.Set(toIncidentKey(incident.ID), incident)
	if err != nil {
		return nil, errors.Wrap(err, "failed to store new incident")
	} else if !saved {
		return nil, errors.New("failed to store new incident")
	}

	// Update Headers
	if err := s.updateHeader(incident); err != nil {
		return nil, errors.Wrap(err, "failed to update headers")
	}

	return incident, nil
}

// UpdateIncident updates an incident.
func (s *StoreImpl) UpdateIncident(incident *Incident) error {
	if incident == nil {
		return errors.New("incident is nil")
	}
	if incident.ID == "" {
		return errors.New("ID should be set")
	}

	headers, err := s.getIDHeaders()
	if err != nil {
		return errors.Wrap(err, "failed to get all headers value")
	}

	if _, exists := headers[incident.ID]; !exists {
		return errors.Errorf("incident with id (%s) does not exist", incident.ID)
	}

	saved, err := s.pluginAPI.KV.Set(toIncidentKey(incident.ID), incident)
	if err != nil {
		return errors.Wrap(err, "failed to update incident")
	} else if !saved {
		return errors.New("failed to update incident")
	}

	// Update Headers
	if err := s.updateHeader(incident); err != nil {
		return errors.Wrap(err, "failed to update headers")
	}

	return nil
}

// GetIncident Gets an incident by ID.
func (s *StoreImpl) GetIncident(id string) (*Incident, error) {
	headers, err := s.getIDHeaders()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all headers value")
	}

	if _, exists := headers[id]; !exists {
		return nil, errors.Errorf("incident with id (%s) does not exist", id)
	}

	var incident Incident
	if err = s.pluginAPI.KV.Get(toIncidentKey(id), &incident); err != nil {
		return nil, errors.Wrap(err, "failed to get incident")
	}

	return &incident, nil
}

// GetAllIncidents Gets all incidents
func (s *StoreImpl) GetAllIncidents() ([]Incident, error) {
	return nil, errors.New("not implemented")
}

// NukeDB Removes all incident related data.
func (s *StoreImpl) NukeDB() error {
	return s.pluginAPI.KV.DeleteAll()
}

// toIncidentKey converts an incident to an internal key used to store in the KV Store.
func toIncidentKey(incidentID string) string {
	return incidentKey + incidentID
}

func toHeader(headers idHeaderMap) []Header {
	var result []Header
	for _, value := range headers {
		result = append(result, value)
	}

	return result
}

func (s *StoreImpl) getIDHeaders() (idHeaderMap, error) {
	headers := idHeaderMap{}
	if err := s.pluginAPI.KV.Get(allHeadersKey, &headers); err != nil {
		return nil, errors.Wrap(err, "failed to get all headers value")
	}
	return headers, nil
}

func (s *StoreImpl) updateHeader(incident *Incident) error {
	headers, err := s.getIDHeaders()
	if err != nil {
		return errors.Wrap(err, "failed to get all headers")
	}

	headers[incident.ID] = incident.Header

	// TODO: Should be using CompareAndSet, but deep copy is expensive.
	if saved, err := s.pluginAPI.KV.Set(allHeadersKey, headers); err != nil {
		return errors.Wrap(err, "failed to set all headers value")
	} else if !saved {
		return errors.New("failed to set all headers value")
	}

	return nil
}