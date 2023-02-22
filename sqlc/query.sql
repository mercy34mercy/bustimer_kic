-- name: GetBusinfoFromDestination :many
SELECT * FROM busstop_url
WHERE destination = ? AND busstop  = ? ;

-- name: GetBusinfoFromBusname :many
SELECT * FROM busstop_url
WHERE busname = ? AND busstop  = ? ;

