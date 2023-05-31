CREATE TABLE comment_like (
    user_id INTEGER,
    comment_id INTEGER,
    PRIMARY KEY (user_id, comment_id)
)