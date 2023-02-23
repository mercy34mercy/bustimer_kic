-- name: GetBusinfoFromDestination :many
SELECT * FROM busstop_urls
WHERE destination = ? AND busstop  = ? ;

-- name: GetBusinfoFromBusname :many
SELECT * FROM busstop_urls
WHERE busname = ? AND busstop  = ? ;

-- name: GetBusstopAndDestination :many
SELECT busstop,destination FROM busstop_urls
WHERE busname = ? ;

-- name: CreateBusstopUrl :one
INSERT INTO busstop_urls (
    busstop,busname,destination,url
) VALUES (
    ?,?,?,?
)
RETURNING *;