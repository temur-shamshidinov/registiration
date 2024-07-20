CREATE TABLE users (
    user_id             UUID            NOT NULL    PRIMARY KEY,
    username            VARCHAR(30)     NOT NULL    UNIQUE,
    gmail               VARCHAR(60)     NOT NULL    UNIQUE,
    password            VARCHAR(60)     NOT NULL
);