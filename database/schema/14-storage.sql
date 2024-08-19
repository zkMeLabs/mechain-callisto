-- bucket tables
CREATE TABLE bucket (
    id BIGINT PRIMARY KEY,
    bucket_id BYTEA NOT NULL UNIQUE,
    bucket_name TEXT NOT NULL UNIQUE CHECK (
        LENGTH(bucket_name) BETWEEN 3 AND 63
    ),
    owner TEXT NOT NULL,
    payment_address TEXT,
    global_virtual_group_family_id INT NOT NULL,
    operator TEXT,
    source_type TEXT,
    charged_read_quota BIGINT,
    visibility TEXT,
    status TEXT,
    delete_at BIGINT,
    delete_reason TEXT,
    storage_size DECIMAL(65, 0) NOT NULL,
    charge_size DECIMAL(65, 0) NOT NULL,
    create_at BIGINT,
    create_tx_hash TEXT NOT NULL,
    create_time TIMESTAMPTZ,
    update_at BIGINT,
    update_tx_hash TEXT NOT NULL,
    update_time TIMESTAMPTZ,
    removed BOOLEAN DEFAULT FALSE,
    off_chain_status INT NOT NULL DEFAULT 0,
    tags JSONB,
    UNIQUE (bucket_id),
    UNIQUE (bucket_name)
);
CREATE INDEX idx_bucket_owner ON bucket(owner);
CREATE INDEX idx_bucket_gvgf_id ON bucket(global_virtual_group_family_id);
-- object tables
CREATE TABLE object (
    id BIGINT PRIMARY KEY,
    bucket_id BYTEA NOT NULL,
    bucket_name TEXT NOT NULL,
    object_id BYTEA NOT NULL UNIQUE,
    object_name TEXT,
    creator TEXT,
    owner TEXT NOT NULL,
    local_virtual_group_id INT,
    operator TEXT,
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
    create_time TIMESTAMPTZ,
    update_at BIGINT,
    update_tx_hash TEXT NOT NULL,
    sealed_tx_hash TEXT,
    update_time TIMESTAMPTZ,
    removed BOOLEAN DEFAULT FALSE,
    tags JSONB,
    is_updating BOOLEAN DEFAULT FALSE,
    content_updated_time TIMESTAMPTZ,
    updater TEXT,
    version BIGINT,
    CONSTRAINT fk_bucket FOREIGN KEY (bucket_id) REFERENCES bucket(bucket_id)
);
CREATE INDEX idx_object_bucket_id ON object(bucket_id);
CREATE INDEX idx_object_bucket_name_object_name ON object(bucket_name, object_name);
CREATE INDEX idx_object_owner ON object(owner);
CREATE INDEX idx_object_local_virtual_group_id ON object(local_virtual_group_id);
-- group tables
CREATE TABLE storage_group (
    id BIGINT PRIMARY KEY,
    owner BYTEA NOT NULL,
    group_id TEXT NOT NULL,
    group_name TEXT,
    source_type TEXT,
    extra TEXT,
    account_id TEXT NOT NULL,
    operator TEXT,
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
    group_id TEXT NOT NULL,
    member TEXT NOT NULL,
    expiration_time BIGINT,
    CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES storage_group(group_id)
);