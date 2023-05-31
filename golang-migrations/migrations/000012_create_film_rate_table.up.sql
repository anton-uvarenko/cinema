CREATE TABLE film_rate (
    users_id INTEGER,
    film_id INTEGER,
    rating INTEGER,
    PRIMARY KEY (users_id, film_id)
)