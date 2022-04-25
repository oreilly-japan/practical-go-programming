-- name: CreateAuthor :one
INSERT INTO author (id, name) VALUES ($1, $2) RETURNING *;

-- name: ListBookOverPrice :many
SELECT b.title, a.name, b.price
FROM book b LEFT JOIN author a ON 1 = 1 AND b.author_id = a.id
WHERE price > $1
ORDER BYb.title;
