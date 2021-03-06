package usecases

import (
	"testing"

	"team-service/domain"
)

type mockLogger struct {
}

func (logger *mockLogger) Log(message string) error {
	return nil
}

type mockEventHandler struct {

}

// Publish sends a new message to the event bus.
func (ev mockEventHandler) Publish(queueName string, evt domain.Event) error {
	println(string(evt.AsEvent()))

	return nil
}

type mockTeamRepository struct {
}

func (repo *mockTeamRepository) Store(team *domain.Team) string {
	newTeamID := "123"

	return newTeamID
}

func (repo *mockTeamRepository) Update(team *domain.Team) *domain.Team {
	return team
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

func (repo *mockTeamRepository) Search(searchTerm string) []domain.Team {
	gk := &domain.Player{
		Name:     "James Eastham",
		Position: "GK",
	}

	def := &domain.Player{
		Name:     "Harry Eastham",
		Position: "DEF",
	}

	team := domain.Team{
		Name: "My FC",
	}

	team.AddPlayer(gk)
	team.AddPlayer(def)

	teams := make([]domain.Team, 1)
	teams[0] = team

	return teams
}

func TestCanCreateTeam(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	team := &CreateTeamRequest{
		Name: "Cornwall FC",
	}

	createTeamResponse, err := teamInteractor.CreateTeam(team)

	if err != nil || len(createTeamResponse.ID) == 0 {
		t.Fatalf("Team has not been created")
	}
}

func TestCanCreateTeam_EmptyName_ShouldError(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	team := &CreateTeamRequest{
		Name: "",
	}

	createTeamResponse, err := teamInteractor.CreateTeam(team)

	if err == nil || len(createTeamResponse.Errors) == 0 {
		t.Fatalf("Creating a team with no name should throw error")
	}
}

func TestCanRetrievePlayersFromATeam(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	team, err := teamInteractor.FindByID("1")

	if err != nil || len(team.Players) != 2 {
		t.Fatalf("Players not retrieved")
	}
}

func TestCanAddPlayerToATeam(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	addPlayerToTeamReq := &AddPlayerToTeamRequest{
		TeamID:         "1",
		PlayerName:     "James Eastham",
		PlayerPosition: "ST",
	}

	team, err := teamInteractor.AddPlayerToTeam(addPlayerToTeamReq)

	if err == nil || len(team.Players) > 2 {
		t.Fatalf("Player has been added and shouldn't have been")
	}
}

func TestCanAddDuplicatePlayerToATeam(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	addPlayerToTeamReq := &AddPlayerToTeamRequest{
		TeamID:         "1",
		PlayerName:     "Karl Eastham",
		PlayerPosition: "ST",
	}

	team, err := teamInteractor.AddPlayerToTeam(addPlayerToTeamReq)

	if err != nil || len(team.Players) != 3 {
		t.Fatalf("Players not added successfully")
	}
}

func TestCanRemovePlayerFromTeam(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	removePlayerFromTeamReq := &RemovePlayerFromTeamRequest{
		TeamID:         "1",
		PlayerName:     "James Eastham",
		PlayerPosition: "GK",
	}

	response, error := teamInteractor.RemovePlayerFromTeam(removePlayerFromTeamReq)

	if error != nil || len(response.Players) > 1 {
		t.Fatalf("Player should be removed successfully")
	}
}

func TestRemoveInvalidPlayerFromTeam(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	removePlayerFromTeamReq := &RemovePlayerFromTeamRequest{
		TeamID:         "1",
		PlayerName:     "aifnofnwuefiw Eastham",
		PlayerPosition: "GK",
	}

	response, error := teamInteractor.RemovePlayerFromTeam(removePlayerFromTeamReq)

	if error == nil || len(response.Errors) < 1 {
		t.Fatalf("Should throw error")
	}
}

func TestCanLoadAllTeams(t *testing.T) {
	teamInteractor := createInMemTeamInteractor()

	teamList, error := teamInteractor.Search("")

	if len(teamList.Teams) < 1 || error != nil {
		t.Fatalf("Empty search should return one team")
	}
}

func createInMemTeamInteractor() *TeamInteractor {
	teamInteractor := &TeamInteractor{
		TeamRepository: &mockTeamRepository{},
		EventHandler: &mockEventHandler{},
		Logger:         &mockLogger{},
	}

	return teamInteractor
}
