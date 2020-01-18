package main

import (
	"context"

	"team-service/usecases"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds al possible endpoints for the team servive.
type Endpoints struct {
	Create               endpoint.Endpoint
	GetByID              endpoint.Endpoint
	AddPlayerToTeam      endpoint.Endpoint
	RemovePlayerFromTeam endpoint.Endpoint
	Search               endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the team service.
func MakeEndpoints(teamInteractor *usecases.TeamInteractor) Endpoints {
	return Endpoints{
		Create:               createTeamEndpoint(teamInteractor),
		GetByID:              loadTeamEndpoint(teamInteractor),
		AddPlayerToTeam:      addPlayerToTeamEndpoint(teamInteractor),
		RemovePlayerFromTeam: removePlayerFromTeamEndpoint(teamInteractor),
		Search:               searchEndpoint(teamInteractor),
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
	Team *usecases.TeamDTO `json:"team"`
	Err  error             `json:"error,omitempty`
}

func (r loadTeamResponse) error() error { return r.Err }

func loadTeamEndpoint(teamInteractor *usecases.TeamInteractor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loadTeamRequest)
		team, err := teamInteractor.FindByID(req.ID)

		return loadTeamResponse{Team: team, Err: err}, nil
	}
}

type playerManagementRequest struct {
	ID       string
	Name     string
	Position string
	Type     string
}

type playerManagementResponse struct {
	Players []usecases.PlayerDTO `json:"players"`
	Err     error                `json:"error,omitempty`
}

func (r playerManagementResponse) error() error { return r.Err }

func addPlayerToTeamEndpoint(teamInteractor *usecases.TeamInteractor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(playerManagementRequest)
		players, err := teamInteractor.AddPlayerToTeam(&usecases.AddPlayerToTeamRequest{
			TeamID:         req.ID,
			PlayerName:     req.Name,
			PlayerPosition: req.Position,
		})

		return playerManagementResponse{Players: players.Players, Err: err}, nil
	}
}

func removePlayerFromTeamEndpoint(teamInteractor *usecases.TeamInteractor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(playerManagementRequest)
		players, err := teamInteractor.RemovePlayerFromTeam(&usecases.RemovePlayerFromTeamRequest{
			TeamID:         req.ID,
			PlayerName:     req.Name,
			PlayerPosition: req.Position,
		})

		return playerManagementResponse{Players: players.Players, Err: err}, nil
	}
}

type searchTeamResponse struct {
	Teams []usecases.TeamDTO `json:"teams"`
	Err   error              `json:"error,omitempty`
}

func searchEndpoint(teamInteractor *usecases.TeamInteractor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)

		searchResponse, err := teamInteractor.Search(req)

		return searchTeamResponse{Teams: searchResponse.Teams, Err: err}, nil
	}
}
