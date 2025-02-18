CREATE TABLE Tasks(
    id BIGSERIAL PRIMARY KEY,
    machineId BIGINT NOT NULL REFERENCES Machine(id),
    shiftId BIGINT REFERENCES Shifts(id),
    frequency TaskFrequency NOT NULL,
    taskPriority TaskPriority NOT NULL,
    description TEXT NOT NULL,
    createdBy BIGINT NOT NULL REFERENCES Users(id),
    createdAt DATE NOT NULL,
    verifiedBy BIGINT REFERENCES Users(id),
    verifiedAt DATE,
    completedBy BIGINT REFERENCES Users(id),
    completedAt DATE,
    status TaskStatus NOT NULL DEFAULT 'todo',
    comment TEXT,
    movedInProgressBy BIGINT REFERENCES Users(id),
    movedInProgressAt DATE
);