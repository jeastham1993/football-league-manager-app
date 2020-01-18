package infrastructure

import (
	"strings"
	"sync"

	"team-service/domain"

	"github.com/rs/xid"
)

// InMemTeamRepo manages in memory storage of teams.
type InMemTeamRepo struct {
	mtx   sync.RWMutex
	teams map[string]*domain.Team
}

// NewInMemTeamRepo creates a new InMemTeamRepo
func NewInMemTeamRepo() *InMemTeamRepo {
	return &InMemTeamRepo{
		teams: make(map[string]*domain.Team),
	}
}

// Store creates a new team in the in memory storage.
func (r *InMemTeamRepo) Store(team *domain.Team) string {
	team.ID = xid.New().String()
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.teams[team.ID] = team
	return team.ID
}

// FindByID gets a team from the in memory store.
func (r *InMemTeamRepo) FindByID(teamID string) *domain.Team {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.teams[teamID]; ok {
		return val
	}

	return nil
}

// Update updates an existing team record
func (r *InMemTeamRepo) Update(team *domain.Team) *domain.Team {
	r.teams[team.ID] = team
	return team
}

// Search looks for a record in the database with the requiste search term
func (r *InMemTeamRepo) Search(searchTerm string) []domain.Team {
	teamsResponse := make([]domain.Team, 0)
	addedTeamCounter := 0

	for _, t := range r.teams {
		if strings.Contains(t.Name, searchTerm) {
			teamsResponse = append(teamsResponse, domain.Team{
				ID:      t.ID,
				Name:    t.Name,
				Players: t.Players,
			})

			addedTeamCounter++
		}
	}

	return teamsResponse
}
