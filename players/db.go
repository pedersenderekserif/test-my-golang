package players

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pedersenderekserif/test-my-golang/utils"
)

func Get(
	ctx context.Context,
	teamId *uuid.UUID,
) (
	result Players,
	err error,
) {

	sqlSelect := `
SELECT
	id
	, team_id
	, name
	, created_at
	, updated_at
	, deleted_at
FROM players
WHERE deleted_at IS NULL`

	args := []interface{}{}

	var rows pgx.Rows

	if args != nil && len(args) > 0 {
		rows, err = utils.GetDbPool().Query(context.Background(), sqlSelect, args)
	} else {
		rows, err = utils.GetDbPool().Query(context.Background(), sqlSelect)
	}

	defer rows.Close()

	for rows.Next() {

		var p Player

		if err = rows.Scan(
			&p.ID,
			&p.TeamID,
			&p.Name,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %v", err)
		}

		result = append(result, p)
	}

	return result, nil
}
