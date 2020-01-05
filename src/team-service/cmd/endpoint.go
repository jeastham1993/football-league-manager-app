package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds al possible endpoints for the team servive.
type Endpoints struct {
	Create  endpoint.Endpoint
	GetByID endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the team service.
func MakeEndpoints(svc TeamService) Endpoints {
	return Endpoints{
		Create:  createTeamEndpoint(svc),
		GetByID: loadTeamEndpoint(svc),
	}
}

type createTeamRequest struct {
	TeamName string
}

type createTeamResponse struct {
	Team *Team `json:"team,omitempty"`
	Err  error `json:"error,omitempty"`
}

func (r createTeamResponse) error() error { return r.Err }

func createTeamEndpoint(s TeamService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createTeamRequest)
		createdTeam, err := s.CreateTeam(req.TeamName)

		return createTeamResponse{Team: createdTeam, Err: err}, nil
	}
}

type loadTeamRequest struct {
	ID string
}

type loadTeamResponse struct {
	TeamData *Team `json:"team,omitempty"`
	Err      error `json:"error,omitempty`
}

func (r loadTeamResponse) error() error { return r.Err }

func loadTeamEndpoint(s TeamService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loadTeamRequest)
		team, err := s.LoadTeam(req.ID)

		return loadTeamResponse{TeamData: team, Err: err}, nil
	}
}
