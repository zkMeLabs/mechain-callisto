CREATE TABLE object (
    id BIGINT PRIMARY KEY,
    bucket_id BYTEA NOT NULL,
    bucket_name VARCHAR(64) NOT NULL,
    object_id BYTEA NOT NULL UNIQUE,
    object_name VARCHAR(1024),
    creator BYTEA,
    owner BYTEA NOT NULL,
    local_virtual_group_id INT,
    operator BYTEA,
    payload_size BIGINT,
    visibility VARCHAR(50),
    content_type VARCHAR(255),
    status VARCHAR(50),
    redundancy_type VARCHAR(50),
    source_type VARCHAR(50),
    checksums TEXT [],
    delete_at BIGINT,
    delete_reason VARCHAR(256),
    create_at BIGINT,
    create_tx_hash BYTEA NOT NULL,
    create_time TIMESTAMPTZ,
    update_at BIGINT,
    update_tx_hash BYTEA NOT NULL,
    sealed_tx_hash BYTEA,
    update_time TIMESTAMPTZ,
    removed BOOLEAN DEFAULT FALSE,
    tags JSONB,
    is_updating BOOLEAN DEFAULT FALSE,
    content_updated_time TIMESTAMPTZ,
    updater BYTEA,
    version BIGINT,
    CONSTRAINT fk_bucket FOREIGN KEY (bucket_id) REFERENCES bucket(bucket_id)
);
CREATE INDEX idx_object_bucket_id ON object(bucket_id);
CREATE INDEX idx_object_bucket_name_object_name ON object(bucket_name, object_name);
CREATE INDEX idx_object_owner ON object(owner);
CREATE INDEX idx_object_local_virtual_group_id ON object(local_virtual_group_id);