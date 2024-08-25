CREATE TABLE objects (
    id SERIAL PRIMARY KEY,
    bucket_id TEXT NOT NULL,
    bucket_name TEXT NOT NULL,
    object_id TEXT NOT NULL UNIQUE,
    object_name TEXT,
    creator_address TEXT,
    owner_address TEXT NOT NULL,
    local_virtual_group_id INT,
    operator_address TEXT,
    payload_size BIGINT,
    visibility TEXT,
    content_type TEXT,
    status TEXT,
    redundancy_type TEXT,
    source_type TEXT,
    checksums TEXT [],
    delete_at BIGINT,
    delete_reason TEXT,
    create_at BIGINT,
    create_tx_hash TEXT NOT NULL,
    create_evm_tx_hash TEXT NOT NULL,
    create_time TIMESTAMPTZ,
    update_at BIGINT,
    update_tx_hash TEXT NOT NULL,
    update_evm_tx_hash TEXT NOT NULL,
    sealed_tx_hash TEXT,
    sealed_evm_tx_hash TEXT NOT NULL,
    update_time TIMESTAMPTZ,
    removed BOOLEAN DEFAULT FALSE,
    tags JSONB,
    is_updating BOOLEAN DEFAULT FALSE,
    content_updated_time TIMESTAMPTZ,
    updater TEXT,
    version BIGINT,
    CONSTRAINT fk_bucket FOREIGN KEY (bucket_id) REFERENCES buckets(bucket_id)
);
CREATE INDEX idx_object_bucket_id ON objects(bucket_id);
CREATE INDEX idx_object_bucket_name_object_name ON objects(bucket_name, object_name);
CREATE INDEX idx_object_owner ON objects(owner_address);
CREATE INDEX idx_object_local_virtual_group_id ON objects(local_virtual_group_id);