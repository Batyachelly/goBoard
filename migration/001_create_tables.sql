CREATE TABLE board (
    id serial PRIMARY KEY,
    status INT NOT NULL,
    title VARCHAR (255) NOT NULL
);

CREATE TABLE message (
    id serial PRIMARY KEY,
    thread_id SERIAL NOT NULL ,
    board_id INT NOT NULL,
    status INT NOT NULL,
    title VARCHAR (255) NOT NULL,
    text TEXT NOT NULL,
    content VARCHAR(255) NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (board_id) REFERENCES board (id)
);
---- create above / drop below ----
DROP TABLE board CASCADE;
DROP TABLE message;
