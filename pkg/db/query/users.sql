-- name: CreateUser :one
INSERT INTO users(email,username,password)
  VALUES ($1,$2,$3) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE ID=$1;

-- name: GetAllUsers :many
SELECT * FROM users LIMIT COALESCE($1,10);