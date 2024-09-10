-- +goose Up

CREATE TABLE following_entries (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    entry_id UUID NOT NULL REFERENCES entries(id) ON DELETE CASCADE,
    UNIQUE(user_id, entry_id)
);

-- +goose Down
DROP TABLE following_entries;