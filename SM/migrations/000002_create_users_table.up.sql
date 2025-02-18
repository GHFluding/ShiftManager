CREATE TABLE Users(
    id BIGSERIAL PRIMARY KEY,
    bitrixid BIGINT NOT NULL,
    name TEXT NOT NULL,
    role UserRole NOT NULL
);