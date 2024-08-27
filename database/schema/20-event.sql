CREATE TABLE bucket_events (
    id SERIAL PRIMARY KEY,
    bucket_id TEXT NOT NULL,
    height BIGINT NOT NULL,
    tx_hash TEXT NOT NULL,
    evm_tx_hash TEXT NOT NULL,
    event TEXT NOT NULL
);
CREATE INDEX idx_bucket_event_bucket_id ON bucket_events(bucket_id);
CREATE INDEX idx_bucket_event_tx_hash ON bucket_events (tx_hash);
CREATE INDEX idx_bucket_event_evm_tx_hash ON bucket_events (evm_tx_hash);
CREATE TABLE object_events (
    id SERIAL PRIMARY KEY,
    object_id TEXT NOT NULL,
    height BIGINT NOT NULL,
    tx_hash TEXT NOT NULL,
    evm_tx_hash TEXT NOT NULL,
    event TEXT NOT NULL
);
CREATE INDEX idx_object_event_object_id ON object_events(object_id);
CREATE INDEX idx_object_event_tx_hash ON object_events (tx_hash);
CREATE INDEX idx_object_event_evm_tx_hash ON object_events (evm_tx_hash);
CREATE TABLE group_events (
    id SERIAL PRIMARY KEY,
    group_id TEXT NOT NULL,
    height BIGINT NOT NULL,
    tx_hash TEXT NOT NULL,
    evm_tx_hash TEXT NOT NULL,
    event TEXT NOT NULL
);
CREATE INDEX idx_group_event_group_id ON group_events(group_id);
CREATE INDEX idx_group_event_tx_hash ON group_events (tx_hash);
CREATE INDEX idx_group_event_evm_tx_hash ON group_events (evm_tx_hash);
CREATE TABLE sp_events (
    id SERIAL PRIMARY KEY,
    sp_id TEXT NOT NULL,
    height BIGINT NOT NULL,
    tx_hash TEXT NOT NULL,
    evm_tx_hash TEXT NOT NULL,
    event TEXT NOT NULL
);
CREATE INDEX idx_sp_event_sp_id ON sp_events(sp_id);
CREATE INDEX idx_sp_event_tx_hash ON sp_events (tx_hash);
CREATE INDEX idx_sp_event_evm_tx_hash ON sp_events (evm_tx_hash);