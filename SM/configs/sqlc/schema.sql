CREATE TYPE UserRole AS ENUM ('engineer', 'worker', 'master', 'manager', 'admin');
CREATE TYPE TaskPriority AS ENUM ('low', 'middle', 'hot_task', 'hot_fix');
CREATE TYPE TaskFrequency AS ENUM ('one_time', 'daily', 'weekly', 'monthly','quarterly');
CREATE TYPE TaskStatus AS ENUM ('todo', 'inProgress', 'failed', 'completed', 'verified');

CREATE TABLE Users(
    id BIGSERIAL PRIMARY KEY,
    bitrixid BIGINT NOT NULL,
    name TEXT NOT NULL,
    role UserRole NOT NULL
);

CREATE TABLE Machine(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    isRepairRequired BOOLEAN DEFAULT FALSE, 
    isActive BOOLEAN DEFAULT TRUE
);

CREATE TABLE Shifts(
    id BIGSERIAL PRIMARY KEY,
    machineId BIGINT NOT NULL REFERENCES Machine(id),
    shift_master BIGINT NOT NULL REFERENCES Users(id),
    createdAt DATE NOT NULL,
    isActive BOOLEAN DEFAULT TRUE,
    deactivatedAt DATE
);

CREATE TABLE Shift_workers(
    shiftId BIGINT NOT NULL REFERENCES Shifts(id),
    userId BIGINT NOT NULL REFERENCES Users(id),
    PRIMARY KEY (shiftId, userId)
);

CREATE TABLE Shift_tasks(
    shiftId BIGINT NOT NULL REFERENCES Shifts(id),
    taskId BIGINT NOT NULL REFERENCES Tasks(id),
    PRIMARY KEY (shiftId, taskId)
);

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