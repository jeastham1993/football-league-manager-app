package main

import (
	uuid "github.com/satori/go.uuid"
)

// Team is the central class in the domain model.
type Team struct {
	ID   string `json:"id"`
	Name string `json:"teamName"`
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
		ID:   newID,
		Name: name,
	}
}
