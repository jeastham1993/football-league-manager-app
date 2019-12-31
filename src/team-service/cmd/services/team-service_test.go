package main

import (
	"testing"
)

type mockTeamDatabaseRepository struct{}

func (m *mockTeamDatabaseRepository) Create(teamName string) string {
	return "OK"
}

func TestCreateTeam_ShouldReturnOk(t *testing.T) {
	databaseRepo := &mockTeamDatabaseRepository{}

	expectedResult := "OK"

	result, err := createTeam(databaseRepo, "my team name")

	if err != nil || expectedResult != result {
		t.Fatalf("Expected %s but got %s", expectedResult, result)
	}
}

func TestCreateTeam_ShouldFailOnEmptyTeam(t *testing.T) {
	databaseRepo := &mockTeamDatabaseRepository{}

	result, err := createTeam(databaseRepo, "")

	if err == nil || len(result) > 0 {
		t.Fatalf("Expected an error but no error was thrown")
	}
}
