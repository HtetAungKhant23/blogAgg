-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at, api_key)
VALUES ($1, $2, $3, $4, uuid_generate_v4()::string)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;
