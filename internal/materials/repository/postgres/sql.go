package db

const queryGetByID string = `SELECT m.uuid, m.type, m.status_id, s.status, m.title, m.content, m.created_at, m.updated_at
FROM materials m
JOIN material_statuses s ON m.status_id = s.id
WHERE m.uuid = $1`
