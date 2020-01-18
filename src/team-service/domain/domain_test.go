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

func TestCanRemoveLastPlayerFromTeam(t *testing.T) {
	team := &Team{}

	team.AddPlayer(&Player{
		Name:     "James",
		Position: "GK",
	})

	removeError := team.RemovePlayer("James", "GK")

	if removeError != nil || len(team.Players) > 0 {
		t.Fatalf("Player should be removed without error")
	}
}

func TestCanRemovePlayerFromTeam(t *testing.T) {
	team := &Team{}

	team.AddPlayer(&Player{
		Name:     "James",
		Position: "GK",
	})

	team.AddPlayer(&Player{
		Name:     "Harry",
		Position: "ST",
	})

	removeError := team.RemovePlayer("James", "GK")

	if removeError != nil || len(team.Players) > 1 {
		t.Fatalf("Player should be removed without error")
	}
}

func TestCanRemoveNonExistentPlayer(t *testing.T) {
	team := &Team{}

	team.AddPlayer(&Player{
		Name:     "James",
		Position: "GK",
	})

	team.AddPlayer(&Player{
		Name:     "Harry",
		Position: "ST",
	})

	removeError := team.RemovePlayer("Karl", "DEF")

	if removeError == nil || removeError.Error() != "Player does not exist and cannot be removed" || len(team.Players) != 2 {
		t.Fatalf("Error should be returned when removing a non existent player")
	}
}
