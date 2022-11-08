CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE weights
(
    id         uuid PRIMARY KEY         NOT NULL DEFAULT uuid_generate_v4(),
    user_id    uuid                     NOT NULL,
    name       text                     NOT NULL,
    weight     float4                   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
