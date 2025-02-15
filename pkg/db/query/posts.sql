-- name: CreatePost :one
INSERT INTO posts(title,content,user_id)
  VALUES ($1,$2,$3) RETURNING *;

-- name: GetPostByID :one
SELECT * FROM posts WHERE ID=$1;

-- name: GetAllPosts :many
SELECT * FROM posts LIMIT COALESCE($1,1000);

-- name: GetPostByUserID :many
SELECT * FROM posts WHERE user_id=$1;