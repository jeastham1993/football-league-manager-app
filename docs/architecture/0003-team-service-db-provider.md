# 2. Database Provider for Team Service datastore

Date: 2019-12-29

## Status

Accepted

## Context

We need to decide which type of database to use for holding team data.

## Decision

The team service will use Amazon DynamoDB as it's data store. Functionally, a relational database may be a better option however the aims of this project or to trial new and different technologies.

## Consequences

Using a NoSQL based database gives a much more dynamic database structure. Adding an extra property to an extisting stored record beocmes trivial, as opose to relational databases where schema changes are required.

That said, for complicated objects, relational databases can work better. We may encounter this issue when storing team and player data respectively.
