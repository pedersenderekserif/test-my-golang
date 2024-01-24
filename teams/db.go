package teams

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pedersenderekserif/test-my-golang/utils"
)

func Get(
	ctx context.Context,
	id *uuid.UUID,
) (
	result Teams,
	err error,
) {

	sqlSelect := `
SELECT
	id
	, name
	, created_at
	updated_at
	, deleted_at
FROM teams
WHERE deleted_at IS NULL`

	args := []interface{}{}

	if id != nil && *id != uuid.Nil {

		args = append(args, id)

		sqlSelect += `
	AND id = $1`
	}

	var rows pgx.Rows

	if args != nil && len(args) > 0 {
		rows, err = utils.GetDbPool().Query(context.Background(), sqlSelect, args)
	} else {
		rows, err = utils.GetDbPool().Query(context.Background(), sqlSelect)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var t Team

		if err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %v", err)
		}

		result = append(result, t)
	}

	return result, nil
}
