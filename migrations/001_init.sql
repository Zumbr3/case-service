-- id gerado na aplicação (Go, uuid.NewV6()), não no banco: o contrato de eventos
-- (docs/architecture/contracts/eventos.md) exige UUID v6 em todo id do sistema

CREATE TABLE cases
(
    id              UUID PRIMARY KEY,
    transaction_id  UUID          NOT NULL UNIQUE,
    account_id      UUID          NOT NULL,
    score           NUMERIC(5, 2) NOT NULL,
    origin          TEXT          NOT NULL CHECK (origin IN ('manual_review', 'auto_block')),
    status          TEXT          NOT NULL DEFAULT 'held' CHECK (status IN ('held', 'released', 'reversed')),
    triggered_rules JSONB         NOT NULL DEFAULT '[]',
    sla_deadline    TIMESTAMP     NOT NULL,
    analyst_id      UUID,
    decision_notes  TEXT,
    decided_at      TIMESTAMP,
    created_at      TIMESTAMP    NOT NULL DEFAULT now()
);

CREATE INDEX idx_cases_status_sla ON cases(status, sla_deadline) WHERE status = 'held';

CREATE TABLE dlq_entries (
    event_id UUID PRIMARY KEY,
    transaction_id UUID,
    error_reason TEXT NOT NULL,
    retry_count INTEGER NOT NULL,
    failed_at TIMESTAMP NOT NULL,
    reprocessed_at TIMESTAMP
);

CREATE TABLE outbox_events (
    id BIGSERIAL PRIMARY KEY,
    aggregate_id UUID NOT NULL,
    topic TEXT NOT NULL,
    payload JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    published_at TIMESTAMP
);