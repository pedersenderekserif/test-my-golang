package teams

import (
	"context"

	"github.com/google/uuid"
)

var TEAMS_MAP map[uuid.UUID]Team

func SetMap(
	ctx context.Context,
) (
	map[uuid.UUID]Team,
	error,
) {

	entries, err := Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, v := range entries {
		TEAMS_MAP[v.ID] = v
	}

	return TEAMS_MAP, nil
}
