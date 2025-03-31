-- name: CreateAlert :one
INSERT INTO alerts (asset_id, severity)
VALUES ($1, $2)
RETURNING *;

-- name: GetAlerts :many
SELECT * FROM alerts
WHERE 
    asset_id = COALESCE($1, asset_id) AND
    severity = COALESCE($2, severity)
ORDER BY id DESC;