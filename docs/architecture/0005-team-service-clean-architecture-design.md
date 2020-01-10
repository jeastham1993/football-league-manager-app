# 5. Clean Architecture Design for the Team Service

Date: 2020-01-10

## Status

Accepted

## Context

Decide the starting point of clean architecture for the team service.

## Decision

### Entities
- Team entity
- Player entity

### Use Cases
- Create a new team
- Update a team
- Delete a team
- Add a player to a team
- Update a players information
- Remove a player from a team

### Interfaces
- Web services for team/player management
- Repositories for team persistance

### Infrastructure
- The database
- Code that handles DB connections
- The HTTP server
- Go kit setup

## Consequences

Go Kit is a ready built and tested set of tools for creating microservice implementations using Go. They have samples for booking applications amongst others and it is well documented.
