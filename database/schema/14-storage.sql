-- storage provider tables
CREATE TABLE storage_provider (
    id BIGINT PRIMARY KEY,
    sp_id INT NOT NULL,
    operator_address TEXT NOT NULL,
    funding_address TEXT NOT NULL,
    seal_address TEXT NOT NULL,
    approval_address TEXT NOT NULL,
    gc_address TEXT NOT NULL,
    total_deposit NUMERIC,
    status TEXT,
    endpoint TEXT,
    moniker TEXT,
    identity TEXT,
    website TEXT,
    security_contact TEXT,
    details TEXT,
    bls_key TEXT,
    update_time_sec BIGINT,
    read_price NUMERIC,
    free_read_quota BIGINT,
    store_price NUMERIC,
    create_at BIGINT,
    create_tx_hash TEXT NOT NULL,
    update_at BIGINT,
    update_tx_hash TEXT NOT NULL,
    removed BOOLEAN DEFAULT FALSE,
    UNIQUE (sp_id)
);
CREATE INDEX idx_sp_id ON storage_provider (sp_id);
CREATE INDEX idx_operator_address ON storage_provider (operator_address);
-- global_virtual_group  table
CREATE TABLE global_virtual_group (
    id BIGINT PRIMARY KEY,
    global_virtual_group_id INT NOT NULL,
    family_id INT NOT NULL,
    primary_sp_id INT NOT NULL,
    secondary_sp_ids TEXT,
    stored_size BIGINT NOT NULL,
    virtual_payment_address TEXT,
    total_deposit NUMERIC,
    create_at BIGINT NOT NULL,
    create_tx_hash TEXT NOT NULL,
    create_time TIMESTAMPTZ NOT NULL,
    update_at BIGINT NOT NULL,
    update_tx_hash TEXT NOT NULL,
    update_time TIMESTAMPTZ NOT NULL,
    removed BOOLEAN DEFAULT FALSE,
    UNIQUE (global_virtual_group_id),
    CONSTRAINT fk_spid FOREIGN KEY (primary_sp_id) REFERENCES storage_provider(sp_id)
);
CREATE INDEX idx_primary_sp_id ON global_virtual_group (primary_sp_id);
-- bucket tables
CREATE TABLE bucket (
    id BIGINT PRIMARY KEY,
    bucket_id TEXT NOT NULL UNIQUE,
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
    bucket_id TEXT NOT NULL,
    bucket_name TEXT NOT NULL,
    object_id TEXT NOT NULL UNIQUE,
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
    owner TEXT NOT NULL,
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
    resource_id TEXT NOT NULL,
    policy_id TEXT NOT NULL,
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