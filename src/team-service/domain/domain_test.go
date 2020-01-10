package domain

import (
	"testing"
)

func TestCanAddValidPlayerToTeam(t *testing.T) {
	team := &Team{}

	team.AddPlayer(&Player{
		Name:     "James",
		Position: "GK",
	})

	if len(team.Players) < 1 {
		t.Fatalf("Player not added")
	}
}

func TestCanAddValidPlayerToTeam_EmptyName_ShouldThrowError(t *testing.T) {
	team := &Team{}

	error := team.AddPlayer(&Player{
		Name:     "",
		Position: "GK",
	})

	if error == nil {
		t.Fatalf("Method should throw error when name is empty")
	}
}

func TestCanAddValidPlayerToTeam_DuplicatePlayer_ShouldThrowError(t *testing.T) {
	team := &Team{}

	firstPlayerAddResult := team.AddPlayer(&Player{
		Name:     "James Eastham",
		Position: "GK",
	})

	secondPlayerAddResult := team.AddPlayer(&Player{
		Name:     "James Eastham",
		Position: "GK",
	})

	if firstPlayerAddResult != nil || secondPlayerAddResult == nil {
		t.Fatalf("Second add of the same name should throw an error")
	}
}

func TestCanAddValidPlayerToTeam_EmptyPosition_ShouldThrowError(t *testing.T) {
	team := &Team{}

	error := team.AddPlayer(&Player{
		Name:     "James",
		Position: "",
	})

	if error == nil {
		t.Fatalf("Method should throw error when position is empty")
	}
}
func TestCanAddValidPlayerToTeam_InvalidPosition_ShouldThrowError(t *testing.T) {
	team := &Team{}

	error := team.AddPlayer(&Player{
		Name:     "James",
		Position: "Prop",
	})

	if error == nil {
		t.Fatalf("Method should throw error when position is empty")
	}
}
