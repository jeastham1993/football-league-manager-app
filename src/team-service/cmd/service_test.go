package main

import (
	"testing"
)

type mockTeamService struct {
}

func NewMockService() TeamService {
	return &mockTeamService{}
}

func (s *mockTeamService) CreateTeam(name string) (*Team, error) {
	return NewTeam(name), nil
}

func (s *mockTeamService) LoadTeam(name string) (*Team, error) {
	return NewTeam("mocked team"), nil
}

func TestCanCreateTeamService(t *testing.T) {
	teamSvc := createInMemTeamSvc()

	if teamSvc == nil {
		t.Fatalf("Team service not created")
	}
}

func TestCanCreateTeam(t *testing.T) {
	teamSvc := createInMemTeamSvc()

	createdTeam, err := teamSvc.CreateTeam("1234")

	if err != nil || createdTeam == nil || len(createdTeam.ID) == 0 {
		t.Fatalf("Team creation failed")
	}
}

func TestCreateTeamFailsOnEmptyName(t *testing.T) {
	teamSvc := createInMemTeamSvc()

	createdTeam, err := teamSvc.CreateTeam("")

	if createdTeam != nil && err == nil {
		t.Fatalf("Error should be null with empty name")
	}
}

func TestLoadTeamShouldReturnInvalidArgument(t *testing.T) {
	teamSvc := createInMemTeamSvc()

	loadedTeams, err := teamSvc.LoadTeam("123")

	if loadedTeams != nil || err != ErrInvalidArgument {
		t.Fatalf("Error should be invalid argument")
	}
}

func TestLoadTeamShouldReturnOneTeam(t *testing.T) {
	teamSvc := createInMemTeamSvc()

	createdTeam, createTeamErr := teamSvc.CreateTeam("mytestteam")

	if createTeamErr == nil {
		loadedTeam, err := teamSvc.LoadTeam(createdTeam.ID)

		if err != nil && loadedTeam == nil {
			t.Fatalf("Team should be populated")
		}
	}
}

func createInMemTeamSvc() TeamService {
	inMemRepo := NewInMemTeamRepository()

	teamSvc := NewService(inMemRepo)

	return teamSvc
}
