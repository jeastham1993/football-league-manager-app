# 6. Event Bus Design

Date: 2020-01-26

## Status

Accepted

## Context

Decide the best implementation of an AWS Event Bus for inter service notifications

## Decision

Use a combination of [SNS](https://aws.amazon.com/sns/) and [SQS](https://aws.amazon.com/sqs/) to manage the notification of events between services.

A service can publish to a logically named topic, and each interested service can create an SQS queue and subscribe that queue to the SNS topic.

## Consequences

Completely decoupled implementation. Service publishes to a notification service and does nothing else thereafter.

A lot of complexity is added, as each interested service needs to manage it's own SQS queue and subscription to the correct topics.
