-- name: InsertLink :exec
INSERT INTO shortlink(
  shortlink,
  url
)
VALUES(
  $1,
  $2
)
RETURNING *;
-- name: GetUrl :one
SELECT
  url
FROM
  shortlink
WHERE
  shortlink = $1;
-- name: DeleteUrl :exec
DELETE FROM shortlink
WHERE shortlink =$1;
