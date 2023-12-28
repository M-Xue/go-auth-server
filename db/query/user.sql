-- name: GetUserByEmail :one
SELECT id, username, email, password, first_name, last_name
FROM Users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id, username, email, password, first_name, last_name 
FROM Users 
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO Users (
	username,
	email,
	password,
	first_name,
	last_name
) VALUES ( $1,$2,$3,$4,$5 )
RETURNING id, username, email, first_name, last_name;

-- name: DeleteUserByID :exec
DELETE FROM Users
WHERE id = $1;
