-- name: GetFolders :many
SELECT * FROM folder
WHERE chat_id = $1;

-- name: GetFiles :many
SELECT * FROM files
WHERE folder_id = $1;