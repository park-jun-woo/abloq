-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;
