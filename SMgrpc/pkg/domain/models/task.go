package models

type Task struct {
	MachineId    int64
	ShiftId      int64
	Frequency    string
	TaskPriority string
	Description  string
}
