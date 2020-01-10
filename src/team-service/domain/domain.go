package domain

import "errors"

var validPositions = [...]string{
	"GK",
	"DEF",
	"MID",
	"ST",
}

// ErrInvalidArgument is thrown when a method argument is invalid.
var ErrInvalidArgument = errors.New("Invalid argument")

// TeamRepository handles the persistance of teams.
type TeamRepository interface {
	FindByID(id string) *Team
	Store(team Team) *string
}

// PlayerRepository repository handles the persistance of players.
type PlayerRepository interface {
	FindByID(id string) *Player
}

// Team is a base entity.
type Team struct {
	ID      string    `json:"id"`
	Name    string    `json:"teamName"`
	Players []*Player `json:"players"`
}

// Player holds data for all the players that a team has.
type Player struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

// AddPlayer adds a player to a team.
func (team *Team) AddPlayer(player *Player) error {
	if len(player.Name) == 0 {
		return ErrInvalidArgument
	}

	for _, v := range team.Players {
		if v.Name == player.Name {
			return ErrInvalidArgument
		}
	}

	if len(player.Position) == 0 {
		return ErrInvalidArgument
	}

	isPositionValid := false

	for _, v := range validPositions {
		if v == player.Position {
			isPositionValid = true
			break
		}
	}

	if isPositionValid == false {
		return ErrInvalidArgument
	}

	team.Players = append(team.Players, player)

	return nil
}
