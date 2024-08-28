CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    owner_address TEXT NOT NULL,
    group_id TEXT NOT NULL,
    group_name TEXT,
    source_type TEXT,
    extra TEXT,
    account_address TEXT NOT NULL,
    operator_address TEXT,
    expiration_time TIMESTAMPTZ,
    create_at BIGINT,
    create_time TIMESTAMPTZ,
    create_tx_hash TEXT NOT NULL,
    create_evm_tx_hash TEXT NOT NULL,
    update_at BIGINT,
    update_time TIMESTAMPTZ,
    update_tx_hash TEXT NOT NULL,
    update_evm_tx_hash TEXT NOT NULL,
    removed BOOLEAN DEFAULT FALSE,
    tags JSONB,
    UNIQUE (group_id)
);
CREATE INDEX idx_group_owner ON "groups"(owner_address);
CREATE INDEX idx_group_group_id ON "groups"(group_id);
CREATE INDEX idx_group_group_name ON "groups"(group_name);
-- group_member table
CREATE TABLE group_member (
    id SERIAL PRIMARY KEY,
    group_id TEXT NOT NULL,
    member TEXT NOT NULL,
    expiration_time TIMESTAMPTZ,
    CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups(group_id)
);