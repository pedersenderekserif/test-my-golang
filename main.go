//go:generate swagger generate spec -m -i .swagger/swagger-tags.json -o .swagger/swagger.json

// Package classification Test My Golang
//
// This API provides a basis for Test My Golang
//
//	BasePath: /
//	Version: 1.0.0
//	Contact: Serif Health<engineering@serifhealth.com> https://www.serifhealth.com
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"context"
	"sync"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/pedersenderekserif/test-my-golang/players"
	"github.com/pedersenderekserif/test-my-golang/teams"
	"github.com/pedersenderekserif/test-my-golang/utils"
	"github.com/sirupsen/logrus"
)

func main() {

	if err := utils.InitDbPool(); err != nil {
		logrus.Fatal(err)
	}

	if err := utils.DbMigrations(); err != nil {
		logrus.Fatal(err)
	}

	if _, err := teams.SetMap(context.Background()); err != nil {
		logrus.Fatal(err)
	}

	if _, err := players.SetTeamPlayersMap(context.Background()); err != nil {
		logrus.Fatal(err)
	}

	router := gin.Default()

	// swagger:route GET /teams teams getTeams
	//
	// Gets teams
	//
	//     Responses:
	//       200: getTeams
	router.GET("/teams", teams.Handler)

	// swagger:route GET /players players getPlayers
	//
	// Gets players
	//
	//     Responses:
	//       200: getPlayers
	router.GET("/players", players.Handler)

	router.Use(static.Serve("/swagger", static.LocalFile(".swagger/", false)))

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()

		logrus.Fatal(router.Run(":9090"))
	}()

	wg.Wait()
}
