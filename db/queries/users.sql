-- name: CreateUser :one
    INSERT INTO users (
                       email,
                       hashed_password
    ) VALUES ($1, $2) RETURNING *;

-- name: GetUserByID :one
    SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
    SELECT * from users WHERE email = $1;

-- name: ListUsers :many
    SELECT * FROM users ORDER BY id
    LIMIT $1 OFFSET $2;

-- name: UpdateUserPassword :one
    UPDATE users SET hashed_password = $1
    WHERE id = $2 RETURNING *;

-- name: DeleteUser :exec
    DELETE from USERS WHERE id = $1;