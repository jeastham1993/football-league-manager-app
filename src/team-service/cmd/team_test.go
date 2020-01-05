package main

import (
	"testing"
)

func TestCreateNewTeamIdSet(t *testing.T) {
	result := NewTeam("testteam")

	if len(result.ID) == 0 {
		t.Fatalf("Team ID is empty")
	}
}

func TestCreateNewTeamNameSet(t *testing.T) {
	result := NewTeam("testteam")

	if result.Name != "testteam" {
		t.Fatalf("Team name has not been set")
	}
}

func TestAddPlayerToTeam(t *testing.T) {
	team := NewTeam("My test team")

	team.AddPlayerToTeam("James", "GK")

	if team.Players[0] == nil {
		t.Fatalf("Player has not been added correctly")
	}
}

func TestAddDuplicatePlayerToTeam(t *testing.T) {
	team := NewTeam("My test team")

	team.AddPlayerToTeam("James Eastham", "GK")
	team.AddPlayerToTeam("James Testing", "GK")
	team.AddPlayerToTeam("James Eastham", "GK")

	if len(team.Players) == 3 {
		t.Fatalf("Duplicate player has been added")
	}
}
