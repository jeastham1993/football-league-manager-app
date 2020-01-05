package main

import (
	uuid "github.com/satori/go.uuid"
)

// Team is the central class in the domain model.
type Team struct {
	ID      string    `json:"id"`
	Name    string    `json:"teamName"`
	Players []*Player `json:"players"`
}

// Player holds data for all the players that a team has
type Player struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

// TeamRepository is an interface for the team database interactions
type TeamRepository interface {
	Store(team *Team) (*Team, error)
	Find(id string) (*Team, error)
}

// NewTeam creates a new, named team.
func NewTeam(name string) *Team {
	newID := uuid.Must(uuid.NewV4()).String()
	return &Team{
		ID:      newID,
		Name:    name,
		Players: []*Player{},
	}
}

// AddPlayerToTeam adds a new player to the specified team, if the player exists the existing player is returned.
func (team *Team) AddPlayerToTeam(playerName string, position string) (*Player, error) {

	for _, v := range team.Players {
		if v.Name == playerName {
			return v, nil
		}
	}

	player := &Player{
		Name:     playerName,
		Position: position,
	}

	team.Players = append(team.Players, player)

	return player, nil
}
