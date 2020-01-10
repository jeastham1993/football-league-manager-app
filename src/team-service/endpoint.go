package main

import (
	"context"

	"github.com/jeastham1993/football-league-manager-app/src/team-service/interfaces"
	"github.com/jeastham1993/football-league-manager-app/src/team-service/usecases"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds al possible endpoints for the team servive.
type Endpoints struct {
	Create  endpoint.Endpoint
	GetByID endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the team service.
func MakeEndpoints(teamInteractor interfaces.TeamInteractor) Endpoints {
	return Endpoints{
		Create:  createTeamEndpoint(teamInteractor),
		GetByID: loadTeamEndpoint(teamInteractor),
	}
}

type createTeamRequest struct {
	TeamName string
}

type createTeamResponse struct {
	TeamID string `json:"teamId,omitempty"`
	Err    error  `json:"error,omitempty"`
}

func (r createTeamResponse) error() error { return r.Err }

func createTeamEndpoint(teamInteractor interfaces.TeamInteractor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createTeamRequest)
		team := &usecases.Team{
			Name: req.TeamName,
		}

		createdTeam, err := teamInteractor.CreateTeam(team)

		return createTeamResponse{TeamID: createdTeam, Err: err}, nil
	}
}

type loadTeamRequest struct {
	ID string
}

type loadTeamResponse struct {
	Players []usecases.Player `json:"players,omitempty"`
	Err     error             `json:"error,omitempty`
}

func (r loadTeamResponse) error() error { return r.Err }

func loadTeamEndpoint(teamInteractor interfaces.TeamInteractor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loadTeamRequest)
		team, err := teamInteractor.Players(req.ID)

		return loadTeamResponse{Players: team, Err: err}, nil
	}
}
