package main

import (
	"errors"
)

// ErrInvalidArgument is a return error message if any method variables are not valid.
var ErrInvalidArgument = errors.New("Invalid argument")

// TeamService contains an interface of all available team service methods.
type TeamService interface {
	// Create team creates a new team in the team datastore.
	CreateTeam(name string) (*Team, error)

	// Load a teams data from the database.
	LoadTeam(teamID string) (*Team, error)
}

// service contains the basic properties of a team service.
type service struct {
	teams TeamRepository
}

// NewService creates a team service with the neccesary dependencies.
func NewService(teamsRepo TeamRepository) TeamService {
	return &service{
		teams: teamsRepo,
	}
}

// CreateTeam allows the creation of a new team in the database.
func (s *service) CreateTeam(name string) (*Team, error) {
	if len(name) == 0 {
		return nil, ErrInvalidArgument
	}

	t := NewTeam(name)

	storeResponse, err := s.teams.Store(t)

	return storeResponse, err
}

// LoadTeam allows the loading of a teams metadata from the database.
func (s *service) LoadTeam(teamID string) (*Team, error) {
	if teamID == "" {
		return nil, ErrInvalidArgument
	}

	t, err := s.teams.Find(teamID)

	return t, err
}
