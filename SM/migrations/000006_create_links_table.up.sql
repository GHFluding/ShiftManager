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
