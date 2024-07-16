-- name: GetProduct :one
SELECT * FROM students 
WHERE id=$1;

-- name: UpdateProduct :exec
UPDATE students SET name=$1, subject=$2, body=$3, created_at=$4, updated_at=$5 
WHERE id=$6;

-- name: CreateProduct :one
INSERT INTO students (
  id , name, subject, body, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM students
WHERE id = $1;

-- name: GetProducts :many
SELECT * FROM students; 