-- bucket tables
CREATE TABLE bucket (
    id BIGINT PRIMARY KEY,
    bucket_id INT NOT NULL UNIQUE,
    bucket_name TEXT NOT NULL UNIQUE CHECK (
        LENGTH(bucket_name) BETWEEN 3 AND 63
    ),
    owner_address TEXT NOT NULL,
    payment_address TEXT,
    global_virtual_group_family_id INT NOT NULL,
    operator_address TEXT,
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
    UNIQUE (bucket_name),
    CONSTRAINT fk_family_id FOREIGN KEY (global_virtual_group_family_id) REFERENCES global_virtual_group_family(global_virtual_group_family_id)
);
CREATE INDEX idx_bucket_owner ON bucket(owner_address);
CREATE INDEX idx_bucket_gvgf_id ON bucket(global_virtual_group_family_id);
-- object tables
CREATE TABLE object (
    id BIGINT PRIMARY KEY,
    bucket_id INT NOT NULL,
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
CREATE INDEX idx_object_owner ON object(owner_address);
CREATE INDEX idx_object_local_virtual_group_id ON object(local_virtual_group_id);
-- group tables
CREATE TABLE storage_group (
    id BIGINT PRIMARY KEY,
    owner_address TEXT NOT NULL,
    group_id TEXT NOT NULL,
    group_name TEXT,
    source_type TEXT,
    extra TEXT,
    account_id TEXT NOT NULL,
    operator_address TEXT,
    expiration_time BIGINT,
    create_at BIGINT,
    create_time TIMESTAMPTZ,
    update_at BIGINT,
    update_time TIMESTAMPTZ,
    removed BOOLEAN DEFAULT FALSE,
    tags JSONB,
    UNIQUE (group_id)
);
CREATE INDEX idx_group_owner ON "storage_group"(owner_address);
CREATE INDEX idx_group_group_id ON "storage_group"(group_id);
CREATE INDEX idx_group_group_name ON "storage_group"(group_name);
-- storage_group_member table
CREATE TABLE storage_group_member (
    id NUMERIC PRIMARY KEY,
    group_id TEXT NOT NULL,
    member TEXT NOT NULL,
    expiration_time TIMESTAMPTZ,
    CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES storage_group(group_id)
);
-- permission table
CREATE TABLE permission (
    id BIGINT PRIMARY KEY,
    principal_type INT NOT NULL,
    principal_value TEXT NOT NULL,
    resource_type TEXT NOT NULL,
    resource_id INT NOT NULL,
    policy_id INT NOT NULL,
    create_time TIMESTAMPTZ NOT NULL,
    update_time TIMESTAMPTZ NOT NULL,
    expiration_time TIMESTAMPTZ NOT NULL,
    removed BOOLEAN NOT NULL,
    UNIQUE (
        principal_type,
        principal_value,
        resource_type,
        resource_id
    )
);
CREATE INDEX idx_policy_id ON permission (policy_id);