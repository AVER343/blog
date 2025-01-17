-- name: CreatePost :one
INSERT INTO posts(title,content,user_id)
  VALUES ($1,$2,$3) RETURNING *;

-- name: GetPostByID :one
SELECT * FROM posts WHERE ID=$1;