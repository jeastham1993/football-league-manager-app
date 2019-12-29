# 2. Service Design

Date: 2019-12-29

## Status

Accepted

## Context

We need to decide on the messages the system will use before commencing development.

## Decision

Services will be designed using the message -> message flow -> contexts pattern to ensure logical boundaries are created from the offset.

See [service design docs](../design/design-docs.md) for the designs used.

## Consequences

An over-comittal to the messages early on gives a good solid base from which the application can be built. However, an over comittal to these messages can lean to a rigid development structure. The messages should be fluid and **are not** set in stone.
