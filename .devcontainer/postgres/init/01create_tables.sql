CREATE TABLE article (
    id        SERIAL,
    title       VARCHAR (40) NOT NULL,
    content      VARCHAR(10000) NOT NULL,
    author   VARCHAR(20),
    post_date     TIMESTAMP,
    update_date     TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE comment (
    id        SERIAL,
    article_id INTEGER,
    content      VARCHAR(1000) NOT NULL,
    commenter   VARCHAR(20),
    post_date     TIMESTAMP,
    update_date     TIMESTAMP,
    PRIMARY KEY (id)
);