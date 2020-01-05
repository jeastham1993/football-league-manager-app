package main

import (
	"sync"
)

type teamRepository struct {
	mtx   sync.RWMutex
	teams map[string]*Team
}

func (r *teamRepository) Store(team *Team) (*Team, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.teams[team.ID] = team
	return team, nil
}

func (r *teamRepository) Find(id string) (*Team, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.teams[id]; ok {
		return val, nil
	}
	return nil, ErrInvalidArgument
}

// NewInMemTeamRepository returns a new instance of a in-memory cargo repository.
func NewInMemTeamRepository() TeamRepository {
	return &teamRepository{
		teams: make(map[string]*Team),
	}
}
