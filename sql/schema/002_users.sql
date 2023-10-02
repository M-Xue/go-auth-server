-- +goose Up
ALTER TABLE user
MODIFY COLUMN id varchar(255);