CREATE TABLE IF NOT EXISTS ormus_source.events
(
    id         UUID PRIMARY KEY DEFAULT uuid(),
    key        text,
    value      text,
    created_at timestamp DEFAULT toTimestamp(now()),
    updated_at timestamp NULL
);