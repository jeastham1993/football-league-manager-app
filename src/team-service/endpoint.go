package main

import (
	"context"

	"team-service/usecases"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds al possible endpoints for the team servive.
type Endpoints struct {
	Create          endpoint.Endpoint
	GetByID         endpoint.Endpoint
	AddPlayerToTeam endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the team service.
func MakeEndpoints(teamInteractor *usecases.TeamInteractor) Endpoints {
	return Endpoints{
		Create:          createTeamEndpoint(teamInteractor),
		GetByID:         loadTeamEndpoint(teamInteractor),
		AddPlayerToTeam: addPlayerToTeamEndpoint(teamInteractor),
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

func createTeamEndpoint(teamInteractor *usecases.TeamInteractor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createTeamRequest)
		team := &usecases.CreateTeamRequest{
			Name: req.TeamName,
		}

		createdTeam, err := teamInteractor.CreateTeam(team)

		return createTeamResponse{TeamID: createdTeam.ID, Err: err}, nil
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

func loadTeamEndpoint(teamInteractor *usecases.TeamInteractor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loadTeamRequest)
		team, err := teamInteractor.Players(req.ID)

		return loadTeamResponse{Players: team, Err: err}, nil
	}
}

type addPlayerToTeamRequest struct {
	ID       string
	Name     string
	Position string
}

type addPlayerToTeamResponse struct {
	Players []usecases.Player `json:"players,omitempty"`
	Err     error             `json:"error,omitempty`
}

func (r addPlayerToTeamResponse) error() error { return r.Err }

func addPlayerToTeamEndpoint(teamInteractor *usecases.TeamInteractor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addPlayerToTeamRequest)
		players, err := teamInteractor.AddPlayerToTeam(&usecases.AddPlayerToTeamRequest{
			TeamID:         req.ID,
			PlayerName:     req.Name,
			PlayerPosition: req.Position,
		})

		return addPlayerToTeamResponse{Players: players.Players, Err: err}, nil
	}
}
