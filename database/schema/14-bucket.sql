CREATE TABLE bucket (
    id BIGINT PRIMARY KEY,
    bucket_id BYTEA NOT NULL UNIQUE,
    bucket_name VARCHAR(64) NOT NULL UNIQUE CHECK (
        LENGTH(bucket_name) BETWEEN 3 AND 63
    ),
    owner BYTEA NOT NULL,
    payment_address BYTEA,
    global_virtual_group_family_id INT NOT NULL,
    operator BYTEA,
    source_type VARCHAR(50),
    charged_read_quota BIGINT,
    visibility VARCHAR(50),
    status VARCHAR(64),
    delete_at BIGINT,
    delete_reason VARCHAR(256),
    storage_size DECIMAL(65, 0) NOT NULL,
    charge_size DECIMAL(65, 0) NOT NULL,
    create_at BIGINT,
    create_tx_hash BYTEA NOT NULL,
    create_time TIMESTAMPTZ,
    update_at BIGINT,
    update_tx_hash BYTEA NOT NULL,
    update_time TIMESTAMPTZ,
    removed BOOLEAN DEFAULT FALSE,
    off_chain_status INT NOT NULL DEFAULT 0,
    tags JSONB,
    UNIQUE (bucket_id),
    UNIQUE (bucket_name)
);
CREATE INDEX idx_bucket_owner ON bucket(owner);
CREATE INDEX idx_bucket_gvgf_id ON bucket(global_virtual_group_family_id);