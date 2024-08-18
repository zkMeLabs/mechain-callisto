CREATE TABLE storage_group (
    id BIGINT PRIMARY KEY,
    owner BYTEA NOT NULL,
    group_id BYTEA NOT NULL,
    group_name VARCHAR(63),
    source_type VARCHAR(63),
    extra VARCHAR(512),
    account_id BYTEA NOT NULL,
    operator BYTEA,
    expiration_time BIGINT,
    create_at BIGINT,
    create_time TIMESTAMPTZ,
    update_at BIGINT,
    update_time TIMESTAMPTZ,
    removed BOOLEAN DEFAULT FALSE,
    tags JSONB,
    UNIQUE (group_id)
);
CREATE INDEX idx_group_owner ON "storage_group"(owner);
CREATE INDEX idx_group_group_id ON "storage_group"(group_id);
CREATE INDEX idx_group_group_name ON "storage_group"(group_name);
CREATE TABLE storage_group_member (
    id NUMERIC PRIMARY KEY,
    group_id BYTEA NOT NULL,
    member BYTEA NOT NULL,
    expiration_time TIMESTAMPTZ,
    CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES storage_group(group_id)
);