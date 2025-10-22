-- +goose Up
CREATE TABLE urls (
    base_url VARCHAR(255) NOT NULL,
    short_url  VARCHAR(255) NOT NULL,
    PRIMARY KEY (short_url)
);


-- +goose Down
DROP TABLE urls;
