CREATE TABLE ports
(
    id          serial PRIMARY KEY,
    port_id     VARCHAR(8)   NULL,
    name        VARCHAR(255) NULL,
    city        VARCHAR(255) NULL,
    country     VARCHAR(255) NULL,
    alias       VARCHAR(255) NULL,
    regions     VARCHAR(255) NULL,
    coordinates VARCHAR(255) NULL,
    province    VARCHAR(255) NULL,
    timezone    VARCHAR(255) NULL,
    unlocs      VARCHAR(255) NULL,
    code        VARCHAR(255) NULL
)
;
