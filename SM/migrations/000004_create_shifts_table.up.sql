CREATE TABLE Shifts(
    id BIGSERIAL PRIMARY KEY,
    machineId BIGINT NOT NULL REFERENCES Machine(id),
    shift_master BIGINT NOT NULL REFERENCES Users(id),
    createdAt DATE NOT NULL,
    isActive BOOLEAN DEFAULT TRUE,
    deactivatedAt DATE
);