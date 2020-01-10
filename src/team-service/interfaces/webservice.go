package interfaces

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jeastham1993/football-league-manager-app/src/team-service/usecases"
)

// TeamInteractor is the interface for managing teams.
type TeamInteractor interface {
	CreateTeam(team *usecases.Team) (string, error)
	Players(teamID string) ([]usecases.Player, error)
}

// WebServiceHandler holds the code for the WebService implementation of the team-service.
type WebServiceHandler struct {
	TeamInteractor TeamInteractor
}

// CreateTeam is the web endpoint for creating a new team.
func (handler WebServiceHandler) CreateTeam(res http.ResponseWriter, req *http.Request) {
	var teamToCreate usecases.Team

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))

	if err := json.Unmarshal(body, &teamToCreate); err != nil {
		return nil, err
	}

	createdTeamID := handler.TeamInteractor.CreateTeam(teamToCreate)

	io.WriteString(res, fmt.Sprintf("TeamID: %d\n", createdTeamID))
}
