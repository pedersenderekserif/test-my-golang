package players

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

var TEAM_PLAYERS_MAP map[uuid.UUID]map[uuid.UUID]Player

func SetTeamPlayersMap(
	ctx context.Context,

) (
	map[uuid.UUID]map[uuid.UUID]Player,
	error,
) {

	TEAM_PLAYERS_MAP = make(map[uuid.UUID]map[uuid.UUID]Player)

	entries, err := Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}

	for i := 0; i < len(entries); i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for _, v := range entries {
				if _, ok := TEAM_PLAYERS_MAP[v.TeamID]; !ok {
					TEAM_PLAYERS_MAP[v.TeamID] = make(map[uuid.UUID]Player)
				}

				TEAM_PLAYERS_MAP[v.TeamID] = map[uuid.UUID]Player{
					v.ID: v,
				}
			}

		}()

	}

	wg.Wait()

	return TEAM_PLAYERS_MAP, nil
}
