-- storage provider
CREATE TABLE storage_providers (
    id SERIAL PRIMARY KEY,
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
    create_evm_tx_hash TEXT NOT NULL,
    update_at BIGINT,
    update_tx_hash TEXT NOT NULL,
    update_evm_tx_hash TEXT NOT NULL,
    removed BOOLEAN DEFAULT FALSE,
    UNIQUE (sp_id)
);
CREATE INDEX idx_sp_id ON storage_providers (sp_id);