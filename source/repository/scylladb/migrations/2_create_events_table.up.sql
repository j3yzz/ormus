CREATE TABLE IF NOT EXISTS events
(
    id         UUID PRIMARY KEY,
    key        text,
    value      text,
    created_at timestamp,
    updated_at timestamp,
);