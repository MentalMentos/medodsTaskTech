-- +goose Up
CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id uuid NOT NULL,
    hashed_token TEXT NOT NULL,
    access_jti TEXT NOT NULL,
    ip_address INET,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS refresh_tokens;
