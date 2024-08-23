CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    resource_type TEXT NOT NULL,
    resource_id INT NOT NULL,
    height BIGINT NOT NULL,
    tx_hash TEXT NOT NULL,
    evm_tx_hash TEXT NOT NULL,
    action INT NOT NULL
);
CREATE INDEX idx_event_tx_hash ON events (tx_hash);
CREATE INDEX idx_event_evm_tx_hash ON events (evm_tx_hash);