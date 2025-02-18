package services

import "sm/internal/database/postgres"

type UserDTO struct {
	ID       int64
	Bitrixid int64
	Name     string
	Role     string
}

type ShiftDTO struct {
	ID            int64
	Machineid     int64
	ShiftMaster   int64
	Createdat     string
	Isactive      bool
	Deactivatedat string
}
type TaskDTO struct {
	ID                int64
	Machineid         int64
	Shiftid           int64
	Frequency         string
	Taskpriority      string
	Description       string
	Createdby         int64
	Createdat         string
	Verifiedby        int64
	Verifiedat        string
	Completedby       int64
	Completedat       string
	Status            string
	Comment           string
	Movedinprogressby int64
	Movedinprogressat string
}

type MachineDTO struct {
	ID               int64
	Name             string
	Isrepairrequired bool
	Isactive         bool
}

type ShiftWorkerDTO struct {
	Shiftid int64
	Userid  int64
}
type ShiftTaskDTO struct {
	Shiftid int64
	Taskid  int64
}

type OutputTypes interface {
	MachineDTO | TaskDTO | ShiftDTO | UserDTO | ShiftWorkerDTO | ShiftTaskDTO
}
type InputTypes interface {
	postgres.Machine | postgres.User | postgres.Shift | postgres.ShiftWorker | postgres.ShiftTask
}
