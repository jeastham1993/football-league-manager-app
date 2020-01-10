package usecases

import (
	"testing"

	"github.com/jeastham1993/football-league-manager-app/src/team-service/domain"
)

type mockLogger struct {
}

func (logger *mockLogger) Log(message string) error {
	return nil
}

type mockTeamRepository struct {
}

func (repo *mockTeamRepository) Store(team *domain.Team) string {
	newTeamID := "123"

	return newTeamID
}

func (repo *mockTeamRepository) FindByID(id string) *domain.Team {
	gk := &domain.Player{
		Name:     "James Eastham",
		Position: "GK",
	}

	def := &domain.Player{
		Name:     "Harry Eastham",
		Position: "DEF",
	}

	team := &domain.Team{
		Name: "My FC",
	}

	team.AddPlayer(gk)
	team.AddPlayer(def)

	return team
}

func TestCanCreateTeam(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	team := &Team{
		Name: "Cornwall FC",
	}

	createdTeamID, err := teamInteractor.CreateTeam(team)

	if err != nil || len(createdTeamID) == 0 {
		t.Fatalf("Team has not been created")
	}
}

func TestCanCreateTeam_EmptyName_ShouldError(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	team := &Team{
		Name: "",
	}

	createdTeamID, err := teamInteractor.CreateTeam(team)

	if err == nil || len(createdTeamID) > 0 {
		t.Fatalf("Creating a team with no name should throw error")
	}
}

func TestCanRetrievePlayersFromATeam(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	players, err := teamInteractor.Players("1")

	if err != nil || len(players) != 2 {
		t.Fatalf("Players not retrieved")
	}
}

func createInMemTeamInteractor() *TeamInteractor {
	teamInteractor := &TeamInteractor{
		TeamRepository: &mockTeamRepository{},
		Logger:         &mockLogger{},
	}

	return teamInteractor
}
