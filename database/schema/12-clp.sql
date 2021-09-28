
CREATE TABLE remove_liquidity_event
(
    signer     TEXT  NOT NULL,
    external_asset JSONB  NOT NULL,
    w_basis_points INT NOT NULL,
    asymmetry       INT NOT NULL,
    height            BIGINT  NOT NULL,
    timestamp TIMESTAMP WITHOUT TIME ZONE,
);
CREATE INDEX remove_liquidity_event_height_index ON remove_liquidity_event (height);

SELECT create_hypertable('remove_liquidity_event','timestamp');


CREATE TABLE create_pool_event
(
    signer     TEXT  NOT NULL,
    external_asset JSONB  NOT NULL,
    native_asset_amount INT NOT NULL,
    external_asset_amount       INT NOT NULL,
    height            BIGINT  NOT NULL,
    timestamp TIMESTAMP WITHOUT TIME ZONE,
);
CREATE INDEX create_pool_event_height_index ON create_pool_event (height);

SELECT create_hypertable('create_pool_event','timestamp');