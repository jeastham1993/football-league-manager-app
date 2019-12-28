One of my developer goals in 2020 is to broaden my horizons. I've always been a .NET developer, which is great for my day to day life. It's familiar, it's friendly and it keeps my brain ticking over. 

But life is about challenging yourself and learning new things. 

I want to finish off the React course I'm working on currently, Blazor seems a fascinating front end framework and Go/Node/Rust have got the juices flowing from a back end point of view.

Throw Docker and Kubernetes into that mix and you have yourself a very excitable developer.

# So what's the plan?

There are a lot of articles/courses dotted around the place, which are written by some brilliant people and are extremely informative. 

However, I don't think there is an awful lot that covers the complete life cycle of a software project. From inception through design to the code monkeys and then deployed to production. 

So there are five key aims this series of articles is going to cover:

## 1. System Design

Give you, the reader, a real deep dive into the world of systems design. We've all heard the words domain driven design banded around, but never used in the real world.

So, whilst my personal goal is to use new technologies (more on that in a sec), from a reader point of view I want you to see past the languages and frameworks.

Follow the development, pull me up on my design decisions, shout when I'm going wrong. I'm by no means an expert, but think I have enough knowledge for this to be a worthwhile series to somebody just getting started.

## 2. Unfamiliar Technologies

Partly through the company I work for being firmly .NET focused, and partly because I'm a bit of wimp, I never really push the boat out too much from a tech point of view. 

MS SQL, .NET and Angular are all extremely familiar technologies I've been working with for many years. Yet there are so many other frameworks out there that may well be better for certain tasks.

I expect to use:

- Node.JS
- GoLang
- ReactJS
- Blazor
- Dapr

To be clear from the start, I have almost zero real world experience with these technologies. So don't expect this series to be a best practice filled, tutorial standard set of steps.

The languages are a learning journey for me, hopefully the process will be a learning journey for you.

## 3. Cloud Ready

Microservices, Docker and Kubernetes have become the industry standard for deploying applications. Separating monoliths into smaller, independent services is the norm.

Frankly, I don't know how it was ever managed any other way. 

So the component parts of the application need to be cloud ready and deploy able anywhere.

I want to take a quick second here to just detail how I define a micro service. A micro-service should always have a single database, but may not necessarily comprise of a single process. I may have a CRUD REST API for data access, with a worker service running behind it managing some arbitrary database operation.

Two separate processes, but one 'service'. *queue inbound abuse from micro service purists*. 

I am likely to use AWS for the hosting, but the application should be infrastructure agnostic.

## 4. Hands Off Deployments

Production releases should happen from a single code push. No file copying, no manual intervention. 

I want seamless, beautiful pipelines.

This means *test driven development*, integration testing and feature flags galore.

The code itself will reside in a git repository, and I will use Azure Pipelines to manage the testing, packaging and releasing.

## 5. Accountability

This final one isn't directly related to the code, but to keeping myself accountable. Every Sunday, I will publish an article detailing that weeks progress. 

That may be me writing a single line of code and getting angry at GoLang. Or maybe it's a completely new production release that is ready for the world to see. 

Whatever it is, there will be something every Sunday. 

# The Project Itself 

I thought long and hard about what to create. I know myself, so know it needs to be something that interests me. 

It also needs to be complicated enough to keep my little noggin' whirring.

So I've dreamed up a slightly un-realistic scenario that I'm hoping doesn't scare off too many of my American readers.

## Setting the scene

The English Premier League (I'm sorry America, but for this entire series football == soccer) need a new league management system.

They are currently running a .NET framework monolith that is getting unwieldy and painful to work with.

They've heard about the magical world of micro services and drafted in an external team to work on a first beta version.

There list of basic requirements are as follows:

- Team registration: Add teams to a league and allow individual teams to submit their squads for the seasons
- Fixture calculator: Calculate the seasons fixtures which will cover each team playing once a week and all playing each other twice (home and away)
- League table: Track results and the league table (including basic team stats goals for, goals against etc etc.)
- Sponsorships: Add sponsors and manage the distribution of the sponsorships to each team
- Transfers: Cover basic transfers between teams and into/out of the league

They are technology and platform agnostic, they just want to see results.

So that's my challenge. You can follow progress [here](https://github.com/jeastham1993/football-league-manager-app) (give it a star if you like :) ).

If there are any features that you would like to see added, or any specifics you want me to really dive into just drop me a comment at any time.

Now, onward to system design.
