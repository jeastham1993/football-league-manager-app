# 4. Database Provider for Team Service datastore

Date: 2020-01-02

## Status

Accepted

## Context

Decide the simplest way to manage inter service communications.

## Decision

We will use the [GoKit](https://github.com/go-kit/kit) library for fascilitaing inter service communications. 

## Consequences

Go Kit is a ready built and tested set of tools for creating microservice implementations using Go. They have samples for booking applications amongst others and it is well documented.
