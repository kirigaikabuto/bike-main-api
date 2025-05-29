-- migrate:up
CREATE TABLE books
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price float NOT NULL
);

-- migrate:down
drop table books;