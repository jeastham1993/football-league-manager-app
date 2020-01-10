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

Comitting to the tiered design proposed by Bob Martin in the book Clean Architecture gives a SOLID base for building applications. Whether this fits well for a GoLang application remains to be seen. But the proposed structure of an application makes sense as far as decoupling goes.

To read more about tiered architecture and the structure checkout out this [article](https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/).
