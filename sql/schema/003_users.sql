-- +goose Up
ALTER TABLE user
ADD UNIQUE (email);

ALTER TABLE user
ADD UNIQUE (username);
