-- name: CreateUser :one
INSERT INTO users(email,username,password)
  VALUES ($1,$2,$3) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM posts WHERE ID=$1;