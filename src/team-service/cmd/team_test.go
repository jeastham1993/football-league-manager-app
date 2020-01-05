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
