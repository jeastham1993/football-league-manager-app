package usecases

import (
	"errors"

	"team-service/domain"
)

// Logger manages how logs are written.
type Logger interface {
	Log(message string) error
}

// CreateTeamRequest holds a reference to new team data.
type CreateTeamRequest struct {
	Name string
}

// CreateTeamResponse holds a reference to new team data.
type CreateTeamResponse struct {
	ID     string
	Name   string
	Errors []string
}

// AddPlayerToTeamRequest adds a player to a team.
type AddPlayerToTeamRequest struct {
	TeamID         string
	PlayerName     string
	PlayerPosition string
}

// AddPlayerToTeamResponse returns the complete set of players for a given team.
type AddPlayerToTeamResponse struct {
	Players []Player
	Errors  []string
}

// Player holds properties for a player object.
type Player struct {
	Name     string
	Position string
}

// TeamInteractor holds all methods for interacting with teams.
type TeamInteractor struct {
	TeamRepository   domain.TeamRepository
	PlayerRepository domain.PlayerRepository
	Logger           Logger
}

// CreateTeam creates a new team in the database.
func (interactor *TeamInteractor) CreateTeam(team *CreateTeamRequest) (*CreateTeamResponse, error) {
	if len(team.Name) == 0 {
		interactor.Logger.Log("Team name cannot be empty")
		var response = &CreateTeamResponse{
			ID:     "",
			Name:   team.Name,
			Errors: make([]string, 1),
		}

		response.Errors[0] = "Team name cannot be empty"

		return response, errors.New("Team name cannot be empty")
	}

	newTeam := &domain.Team{
		Name: team.Name,
	}

	createdTeamID := interactor.TeamRepository.Store(newTeam)

	return &CreateTeamResponse{
		ID:   createdTeamID,
		Name: newTeam.Name,
	}, nil
}

// Players retrieves a list of players for a given team.
func (interactor *TeamInteractor) Players(teamID string) ([]Player, error) {
	var players []Player

	team := interactor.TeamRepository.FindByID(teamID)

	players = make([]Player, len(team.Players))

	for i, player := range team.Players {
		players[i] = Player{player.Name, player.Position}
	}

	return players, nil
}

// AddPlayerToTeam adds a player to a team.
func (interactor *TeamInteractor) AddPlayerToTeam(request *AddPlayerToTeamRequest) (AddPlayerToTeamResponse, error) {
	var response AddPlayerToTeamResponse

	team := interactor.TeamRepository.FindByID(request.TeamID)

	if team != nil {
		newPlayer := &domain.Player{
			Name:     request.PlayerName,
			Position: request.PlayerPosition,
		}

		err := team.AddPlayer(newPlayer)

		if err != nil {
			return response, err
		}

		response.Players = make([]Player, len(team.Players))

		interactor.TeamRepository.Update(team)

		for i, player := range team.Players {
			response.Players[i] = Player{player.Name, player.Position}
		}
	}

	return response, nil
}
