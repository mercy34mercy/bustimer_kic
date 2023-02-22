-- name: GetBusinfo :one
SELECT * FROM busstop_url
WHERE destination = ? AND busstop  = ? ;


SELECT * FROM busstop_url
WHERE busname = ? AND busstop  = ? ;

