-- Create a new session
-- Parameters: id, account_id, refresh_token, user_agent, client_ip, is_blocked, expires_at
-- Returns: Newly created session
-- name: CreateSession :exec
INSERT INTO sessions (id, account_id, refresh_token, user_agent, client_ip, is_blocked, expires_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
