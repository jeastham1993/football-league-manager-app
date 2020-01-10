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
	Store(team Team)
	FindById(id string) Team
	Update(team Team)
}

// PlayerRepository repository handles the persistance of players.
type PlayerRepository interface {
	Store(player Player)
	FindById(id string) Player
	Update(player Player)
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
