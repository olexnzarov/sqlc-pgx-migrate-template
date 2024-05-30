-- name: GetAll :many
SELECT 
    author.*, 
    COALESCE(AVG(book.rating), 0)::FLOAT AS average_rating 
FROM authors author 
    LEFT JOIN books book ON book.author_id = author.id
GROUP BY author.id;

-- name: Get :one
SELECT * FROM authors WHERE id = $1;

-- name: Create :one
INSERT INTO authors (
    name, 
    description
) VALUES ($1, $2) RETURNING id;