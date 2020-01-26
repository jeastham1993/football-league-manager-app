package usecases

import (
	"encoding/json"
	"errors"

	"team-service/domain"
)

// Logger manages how logs are written.
type Logger interface {
	Log(message string) error
}

// EventBus handles interactions with the application event bus.
type EventBus interface {
	Publish(queue string, event domain.Event) error
}

// ErrTeamNotFound is returned when a team is searched for and not found in the database.
var ErrTeamNotFound = errors.New("Specified team not found")

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
	Players []PlayerDTO
	Errors  []string
}

// RemovePlayerFromTeamRequest removes a player from a team
type RemovePlayerFromTeamRequest struct {
	TeamID         string
	PlayerName     string
	PlayerPosition string
}

// RemovePlayerFromTeamResponse returns the player list after having the player removed along with any errors.
type RemovePlayerFromTeamResponse struct {
	Players []PlayerDTO
	Errors  []string
}

// SearchTeamsResponse allows the searching of teams in the database.
type SearchTeamsResponse struct {
	Teams  []TeamDTO
	Errors []string
}

// TeamDTO holds properties for a team object.
type TeamDTO struct {
	ID      string
	Name    string
	Players []PlayerDTO
}

// PlayerDTO holds properties for a player object.
type PlayerDTO struct {
	Name     string
	Position string
}

// TeamInteractor holds all methods for interacting with teams.
type TeamInteractor struct {
	TeamRepository   domain.TeamRepository
	PlayerRepository domain.PlayerRepository
	Logger           Logger
	EventHandler     EventBus
}

// TeamCreatedEvent is published when a new team is created.
type TeamCreatedEvent struct {
	TeamName string
	TeamID   string
}

// AsEvent returns a string representation of the given object.
func (t TeamCreatedEvent) AsEvent() []byte {
	responseBytes, err := json.Marshal(t)

	if err == nil {
		return responseBytes
	}

	return nil
}

// PlayerAddedEvent is published when a new player is added to a team
type PlayerAddedEvent struct {
	TeamName       string
	TeamID         string
	PlayerName     string
	PlayerPosition string
}

// AsEvent returns a string representation of the given object.
func (p PlayerAddedEvent) AsEvent() []byte {
	responseBytes, err := json.Marshal(p)

	if err == nil {
		return responseBytes
	}

	return nil
}

// PlayerRemovedEvent is published when a new player is added to a team
type PlayerRemovedEvent struct {
	TeamName       string
	TeamID         string
	PlayerName     string
	PlayerPosition string
}

// AsEvent returns a string representation of the given object.
func (p PlayerRemovedEvent) AsEvent() []byte {
	responseBytes, err := json.Marshal(p)

	if err == nil {
		return responseBytes
	}

	return nil
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

	interactor.EventHandler.Publish("leaguemanager-info-newteamcreated", TeamCreatedEvent{
		TeamID:   createdTeamID,
		TeamName: team.Name,
	})

	return &CreateTeamResponse{
		ID:   createdTeamID,
		Name: newTeam.Name,
	}, nil
}

// FindByID retrieves data about a specific team
func (interactor *TeamInteractor) FindByID(teamID string) (*TeamDTO, error) {
	var players []PlayerDTO

	team := interactor.TeamRepository.FindByID(teamID)

	if team == nil {
		return nil, ErrTeamNotFound
	}

	players = make([]PlayerDTO, len(team.Players))

	for i, player := range team.Players {
		players[i] = PlayerDTO{player.Name, player.Position}
	}

	return &TeamDTO{
		Name:    team.Name,
		Players: players,
	}, nil
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

		response.Players = make([]PlayerDTO, len(team.Players))

		interactor.TeamRepository.Update(team)

		for i, player := range team.Players {
			response.Players[i] = PlayerDTO{player.Name, player.Position}
		}

		interactor.EventHandler.Publish("leaguemanager-newplayer", &PlayerAddedEvent{
			TeamName:       team.Name,
			TeamID:         team.ID,
			PlayerName:     request.PlayerName,
			PlayerPosition: request.PlayerPosition,
		})
	}

	return response, nil
}

// RemovePlayerFromTeam adds a player to a team.
func (interactor *TeamInteractor) RemovePlayerFromTeam(request *RemovePlayerFromTeamRequest) (RemovePlayerFromTeamResponse, error) {
	var response RemovePlayerFromTeamResponse

	team := interactor.TeamRepository.FindByID(request.TeamID)

	if team != nil {
		err := team.RemovePlayer(request.PlayerName, request.PlayerPosition)

		if err != nil {
			response.Errors = make([]string, 1)
			response.Errors[0] = err.Error()

			return response, err
		}

		response.Players = make([]PlayerDTO, len(team.Players))

		interactor.TeamRepository.Update(team)

		for i, player := range team.Players {
			response.Players[i] = PlayerDTO{player.Name, player.Position}
		}

		interactor.EventHandler.Publish("leaguemanager-playerremoved", &PlayerRemovedEvent{
			TeamName:       team.Name,
			TeamID:         team.ID,
			PlayerName:     request.PlayerName,
			PlayerPosition: request.PlayerPosition,
		})
	}

	return response, nil
}

// Search allows searching across all teams in the database.
func (interactor *TeamInteractor) Search(searchTerm string) (SearchTeamsResponse, error) {
	var response SearchTeamsResponse

	teams := interactor.TeamRepository.Search(searchTerm)

	if teams == nil {
		response.Errors = make([]string, 1)
		response.Errors[0] = "No results found"

		return response, ErrTeamNotFound
	}

	response.Teams = make([]TeamDTO, len(teams))

	for i, team := range teams {
		teamDTO := TeamDTO{team.ID, team.Name, make([]PlayerDTO, len(team.Players))}

		for _, player := range team.Players {
			teamDTO.Players[i] = PlayerDTO{player.Name, player.Position}
		}

		response.Teams[i] = teamDTO
	}

	return response, nil
}
