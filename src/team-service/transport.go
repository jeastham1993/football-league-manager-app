package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// ErrBadRouting indicates an incorrect route has been requested.
var (
	ErrBadRouting = errors.New("bad routing")
)

// MakeHandler returns a handler for the team service.
func MakeHandler(svcEndpoints Endpoints) http.Handler {
	// set-up router and initialize http endpoints
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/team").Handler(kithttp.NewServer(
		svcEndpoints.Create,
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/team/{id}/players").Handler(kithttp.NewServer(
		svcEndpoints.AddPlayerToTeam,
		decodedAddPlayerToTeamRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/team/{id}/players").Handler(kithttp.NewServer(
		svcEndpoints.RemovePlayerFromTeam,
		decodedRemovePlayerFromTeamRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/team/{id}").Handler(kithttp.NewServer(
		svcEndpoints.GetByID,
		decodedGetByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/team").Handler(kithttp.NewServer(
		svcEndpoints.Search,
		decodedSearchRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var createRequest createTeamRequest

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err := json.Unmarshal(body, &createRequest); err != nil {
		return nil, err
	}
	return createRequest, nil
}

func decodedGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return loadTeamRequest{ID: id}, nil
}

func decodedSearchRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	searchTermQueryParam := r.URL.Query()["search"]

	searchTerm := ""

	if len(searchTermQueryParam) > 0 {
		searchTerm = searchTermQueryParam[0]
	}

	return searchTerm, nil
}

func decodedAddPlayerToTeamRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		return nil, ErrBadRouting
	}

	var addPlayerToTeamRequest playerManagementRequest

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err := json.Unmarshal(body, &addPlayerToTeamRequest); err != nil {
		return nil, err
	}

	addPlayerToTeamRequest.ID = id

	return addPlayerToTeamRequest, nil
}

func decodedRemovePlayerFromTeamRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		return nil, ErrBadRouting
	}

	var removePlayerFromTeamRequest playerManagementRequest

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err := json.Unmarshal(body, &removePlayerFromTeamRequest); err != nil {
		return nil, err
	}

	removePlayerFromTeamRequest.ID = id

	return removePlayerFromTeamRequest, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
