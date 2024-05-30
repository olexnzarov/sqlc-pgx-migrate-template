-- name: GetAll :many
SELECT * FROM books ORDER BY published_at;

-- name: Get :one
SELECT * FROM books WHERE id = $1;

-- name: Create :one
INSERT INTO books (
    name, 
    description, 
    author_id, 
    rating,    
    published_at
) VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: Update :execrows
UPDATE books SET 
    name = $2, 
    description = $3, 
    author_id = $4, 
    rating = $5,    
    published_at = $6
 WHERE id = $1;

-- name: Delete :execrows
DELETE FROM books WHERE id = $1;

