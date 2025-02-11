CREATE UserRole AS ENUM ('engineer', 'worker', 'master', 'manager', 'admin');
CREATE TaskPriority AS ENUM ('low', 'middle', 'hot_task', 'hot_fix');
CREATE TaskFrequency AS ENUM ('one_time', 'daily', 'weekly', 'monthly','quarterly');
CREATE TaskStatus AS ENUM ('todo', 'inProgress', 'failed', 'completed', 'verified');
CREATE TABLE Users(
    id  BIGSERIAL PRIMARY KEY,
    bitrixid BIGSERIAL NOT NULL,
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
    machineId BIGSERIAL NOT NULL,
    shift_master BIGSERIAL NOT NULL,
    createdAt DATE NOT NULL,
    isActive BOOLEAN DEFAULT TRUE,
    deactivatedAt DATE NULLABLE,
);
CREATE TABLE Shift_workers(
    shiftId BIGSERIAL NOT NULL,
    userId BIGSERIAL NOT NULL,
);
CREATE TABLE Shift_tasks(
    shiftId BIGSERIAL NOT NULL,
    taskId BIGSERIAL NOT NULL,
);
CREATE TABLE Tasks(
    id BIGSERIAL PRIMARY KEY,
    machineId BIGSERIAL NOT NULL,
    shiftId BIGSERIAL NOT NULL,
    frequency TaskFrequency NOT NULL,
    taskPriority TaskPriority NOT NULL,
    description TEXT NOT NULL,
    createdBy BIGSERIAL NOT NULL,
    createdAt DATE NOT NULL,
    verifiedBy BIGSERIAL NOT NULL,
    verifiedAt DATE NOT NULL,
    completedBy BIGSERIAL NOT NULL,
    completedAt DATE NOT NULL,
    status TaskStatus NOT NULL,
    comment TEXT NULLABLE,
    movedInProgressBy BIGSERIAL NULLABLE,
    movedInProgressAt  DATE NULLABLE
)
