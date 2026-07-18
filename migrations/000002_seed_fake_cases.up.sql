-- Seed fake cases so case_db can be exercised locally/in demos without depending on
-- the Kafka consumer to populate the `cases` table. Ids are fixed UUIDv6 values (as
-- required by docs/architecture/contracts/eventos.md) so this migration is idempotent
-- to reason about and easy to clean up via the matching .down.sql.

INSERT INTO cases (id, transaction_id, account_id, score, origin, status, triggered_rules, sla_deadline, analyst_id, decision_notes, decided_at, created_at)
VALUES
    -- held, manual_review, within SLA
    ('1f175446-2bd6-6000-8000-d337f6230dde', '1f175446-355f-6680-8001-c7dbe7ea7cf3', '1f175446-3ee8-6d00-8002-b14cb7fd03c2',
     42.50, 'manual_review', 'held',
     '[{"ruleID":"R001","name":"velocity_check","partialScore":20.50,"weight":"medium"},{"ruleID":"R002","name":"device_mismatch","partialScore":22.00,"weight":"medium"}]',
     now() + interval '2 days', NULL, NULL, NULL, now() - interval '1 hour'),

    -- held, auto_block, SLA already breached
    ('1f175446-4872-6380-8003-89945c86a14d', '1f175446-51fb-6a00-8004-7f152f543f00', '1f175446-5b85-6080-8005-f19ab99ed3af',
     88.75, 'auto_block', 'held',
     '[{"ruleID":"R010","name":"blacklisted_device","partialScore":88.75,"weight":"high"}]',
     now() - interval '1 day', NULL, NULL, NULL, now() - interval '3 days'),

    -- held, manual_review, SLA due soon
    ('1f175446-650e-6700-8006-314bc5ad82f0', '1f175446-6e97-6d80-8007-6f9ffffbba3c', '1f175446-7821-6400-8008-8d86c0872a47',
     55.55, 'manual_review', 'held',
     '[{"ruleID":"R003","name":"amount_outlier","partialScore":55.55,"weight":"medium"}]',
     now() + interval '1 hour', NULL, NULL, NULL, now() - interval '30 minutes'),

    -- held, auto_block, low score, far SLA
    ('1f175446-81aa-6a80-8009-e7770e61a5c5', '1f175446-8b34-6100-800a-6b6077e8d784', '1f175446-94bd-6780-800b-831a3c393cef',
     5.00, 'auto_block', 'held',
     '[{"ruleID":"R004","name":"new_account","partialScore":5.00,"weight":"low"}]',
     now() + interval '5 days', NULL, NULL, NULL, now()),

    -- released, manual_review
    ('1f175446-9e46-6e00-800c-a762a3c58f0d', '1f175446-a7d0-6480-800d-9309e4e4e9d5', '1f175446-b159-6b00-800e-7767162988fb',
     15.00, 'manual_review', 'released',
     '[{"ruleID":"R005","name":"velocity_check","partialScore":15.00,"weight":"low"}]',
     now() - interval '4 days', '1f175447-799f-6380-8023-655eabd21061', 'False positive, customer verified via support call', now() - interval '3 days', now() - interval '5 days'),

    -- released, auto_block, high score
    ('1f175446-c46c-6800-8010-a5fe4bbaaa1e', '1f175446-cdf5-6e80-8011-659b7034bd7f', '1f175446-d77f-6500-8012-6bf5bcdac6fa',
     95.20, 'auto_block', 'released',
     '[{"ruleID":"R010","name":"blacklisted_device","partialScore":95.20,"weight":"high"}]',
     now() - interval '6 days', '1f175447-8328-6a00-8024-4514d522c01c', 'Manual KYC re-verification passed', now() - interval '5 days', now() - interval '7 days'),

    -- released, manual_review, max score
    ('1f175446-ea92-6200-8014-09654138a8e9', '1f175446-f41b-6880-8015-a730be41fb8b', '1f175446-fda4-6f00-8016-43b1f866fd79',
     100.00, 'manual_review', 'released',
     '[{"ruleID":"R006","name":"amount_outlier","partialScore":60.00,"weight":"high"},{"ruleID":"R007","name":"device_mismatch","partialScore":40.00,"weight":"medium"}]',
     now() - interval '10 days', '1f175447-799f-6380-8023-655eabd21061', 'Reviewed full transaction history, legitimate purchase', now() - interval '9 days', now() - interval '11 days'),

    -- reversed, manual_review
    ('1f175447-10b7-6c00-8018-71ac68e9f133', '1f175447-1a41-6280-8019-7d95870fa9ac', '1f175447-23ca-6900-801a-1d5467f1932f',
     60.00, 'manual_review', 'reversed',
     '[{"ruleID":"R008","name":"velocity_check","partialScore":60.00,"weight":"high"}]',
     now() - interval '2 days', '1f175447-8cb2-6080-8025-01cc3591d6c5', 'Confirmed fraud, chargeback initiated', now() - interval '1 day', now() - interval '3 days'),

    -- reversed, auto_block
    ('1f175447-36dd-6600-801c-3f5650f8821e', '1f175447-4066-6c80-801d-83aa96538f74', '1f175447-49f0-6300-801e-010f37df895a',
     30.10, 'auto_block', 'reversed',
     '[{"ruleID":"R009","name":"new_account","partialScore":30.10,"weight":"medium"}]',
     now() - interval '8 days', '1f175447-8cb2-6080-8025-01cc3591d6c5', 'Account takeover confirmed by security team', now() - interval '7 days', now() - interval '9 days'),

    -- reversed, auto_block, zero score edge case
    ('1f175447-5d03-6000-8020-f7921e7fb677', '1f175447-668c-6680-8021-99ce03324a8c', '1f175447-7015-6d00-8022-5f5510c0fd32',
     0.00, 'auto_block', 'reversed',
     '[]',
     now() - interval '12 days', '1f175447-8328-6a00-8024-4514d522c01c', 'Rule triggered on stale data, reverted after audit', now() - interval '11 days', now() - interval '13 days');
