# Test My Golang

Test My Golang is the Serif Health code-pairing Golang interview project. It's meant to test one's basic understanding of Golang. 

It's an API that supports teams and players who are assigned to a team.

## Prerequisites

The applicant will need to have [Golang](https://go.dev/doc/install), [golang-migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md), [PostgreSQL](https://www.postgresql.org/download/), [Docker](https://docs.docker.com/engine/install/), and [Docker-Compose](https://docs.docker.com/compose/install/) installed on their machine to be able to run `test-my-golang`.

The applicant will be given a link to the github repository for `test-my-golang` at the start of their interview.

A VSCode launch configuration has been provided for ease of setup, but the applicant is recommended to pick and use whatever IDE they're most productive in.

## Requirements

1. Debug application startup errors
   - There are several bug across the application layers that prevent it from successfuly launching the API router
2. Debug players endpoint functionality
   - A bug has been reported in the players search endpoint where players from different teams are being returned when searching by a specific team id
3. Update players object to include current team name
   - A user has requested to view the name of the team the player is currently assigned to
4. Create an endpoint to trade a player to another team
   - A user now needs the ability to trade a player to another team

## Start Application

First the applicant will need to run `make db-start` or equivalent commands to get the docker db started and loaded with test data. 

Then the applicant should be able to start the application from within their IDE and proceed debugging it.