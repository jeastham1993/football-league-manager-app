package usecases

import (
	"errors"

	"github.com/jeastham1993/football-league-manager-app/src/team-service/domain"
)

// Logger manages how logs are written.
type Logger interface {
	Log(message string) error
}

// Team holds a reference to the team data.
type Team struct {
	Name string
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
func (interactor *TeamInteractor) CreateTeam(team *Team) (string, error) {
	if len(team.Name) == 0 {
		interactor.Logger.Log("Team name cannot be empty")
		return "", errors.New("Team name cannot be empty")
	}

	newTeam := &domain.Team{
		Name: team.Name,
	}

	createdTeamID := interactor.TeamRepository.Store(newTeam)

	return createdTeamID, nil
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
