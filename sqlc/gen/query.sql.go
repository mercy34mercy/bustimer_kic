// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: query.sql

package bustimersqlc

import (
	"context"
)

const getBusinfoFromBusname = `-- name: GetBusinfoFromBusname :many
;

SELECT id, busstop, busname, destination, url FROM busstop_url
WHERE busname = ? AND busstop  = ?
`

type GetBusinfoFromBusnameParams struct {
	Busname string
	Busstop string
}

func (q *Queries) GetBusinfoFromBusname(ctx context.Context, arg GetBusinfoFromBusnameParams) ([]BusstopUrl, error) {
	rows, err := q.db.QueryContext(ctx, getBusinfoFromBusname, arg.Busname, arg.Busstop)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BusstopUrl
	for rows.Next() {
		var i BusstopUrl
		if err := rows.Scan(
			&i.ID,
			&i.Busstop,
			&i.Busname,
			&i.Destination,
			&i.Url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBusinfoFromDestination = `-- name: GetBusinfoFromDestination :many
SELECT id, busstop, busname, destination, url FROM busstop_url
WHERE destination = ? AND busstop  = ?
`

type GetBusinfoFromDestinationParams struct {
	Destination string
	Busstop     string
}

func (q *Queries) GetBusinfoFromDestination(ctx context.Context, arg GetBusinfoFromDestinationParams) ([]BusstopUrl, error) {
	rows, err := q.db.QueryContext(ctx, getBusinfoFromDestination, arg.Destination, arg.Busstop)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BusstopUrl
	for rows.Next() {
		var i BusstopUrl
		if err := rows.Scan(
			&i.ID,
			&i.Busstop,
			&i.Busname,
			&i.Destination,
			&i.Url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
