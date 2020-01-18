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

// ErrNonExistentPlayer is thrown when a player is attempted to be removed and doesn't exist.
var ErrNonExistentPlayer = errors.New("Player does not exist and cannot be removed")

// TeamRepository handles the persistance of teams.
type TeamRepository interface {
	FindByID(id string) *Team
	Store(team *Team) string
	Update(team *Team) *Team
	Search(searchTerm string) []Team
}

// PlayerRepository repository handles the persistance of players.
type PlayerRepository interface {
	FindByID(id string) *Player
}

// Event holds details about the event to be published
type Event interface {
	AsEvent() []byte
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

// RemovePlayer removes the given player from the team.
func (team *Team) RemovePlayer(name, position string) error {
	playerRemoved := false

	for i, v := range team.Players {
		if v.Name == name && v.Position == position {
			playerRemoved = true

			if len(team.Players) == 1 {
				team.Players = nil
				break
			}

			team.Players = append(team.Players[:i], team.Players[i+1])
		}
	}

	if playerRemoved {
		return nil
	}

	return ErrNonExistentPlayer
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
