CREATE TABLE comments (
    id SERIAL PRIMARY KEY ,
    text TEXT,
    user_id INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    comment_type VARCHAR(10),
    reply_to INT REFERENCES comments(id),
    rating INT,
    was_edited BOOL,
    film_id INT
)