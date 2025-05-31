package services

import "github.com/GHFluding/ShiftManager/SM/internal/database/postgres"

type User struct {
	ID       int64
	Bitrixid int64
	Name     string
	Role     string
}

func convertUserDB(inp postgres.User) User {
	var out User
	out.ID = inp.ID
	out.Bitrixid = inp.Bitrixid.Int64
	out.Name = inp.Name
	out.Role = string(inp.Name)
	return out
}

type Shift struct {
	ID            int64
	Machineid     int64
	ShiftMaster   int64
	Createdat     string
	Isactive      bool
	Deactivatedat string
}

func convertShiftDB(inp postgres.Shift) Shift {
	var out Shift
	out.ID = inp.ID
	out.Machineid = inp.Machineid
	out.ShiftMaster = inp.ShiftMaster
	out.Createdat = inp.Createdat.Time.String()
	out.Isactive = inp.Isactive.Bool
	if inp.Deactivatedat.Valid {
		out.Deactivatedat = inp.Deactivatedat.Time.String()
	}
	return out
}

type Task struct {
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

func convertTaskDB(inp postgres.Task) Task {
	var out Task
	out.ID = inp.ID
	out.Machineid = inp.Machineid
	if inp.Shiftid.Valid {
		out.Shiftid = inp.Shiftid.Int64
	}
	out.Frequency = string(inp.Frequency)
	out.Taskpriority = string(inp.Taskpriority)
	out.Description = inp.Description
	out.Createdby = inp.Createdby
	out.Createdat = inp.Createdat.Time.String()
	if inp.Verifiedby.Valid {
		out.Verifiedby = inp.Verifiedby.Int64
	}
	if inp.Verifiedat.Valid {
		out.Verifiedat = inp.Verifiedat.Time.String()
	}
	if inp.Completedby.Valid {
		out.Completedby = inp.Completedby.Int64
	}

	if inp.Completedat.Valid {
		out.Completedat = inp.Completedat.Time.String()
	}
	out.Status = string(inp.Status)
	if inp.Comment.Valid {
		out.Comment = inp.Comment.String
	}
	if inp.Movedinprogressby.Valid {
		out.Movedinprogressby = inp.Movedinprogressby.Int64
	}
	if inp.Movedinprogressat.Valid {
		out.Movedinprogressat = inp.Movedinprogressat.Time.String()
	}
	return out
}

type Machine struct {
	ID               int64
	Name             string
	Isrepairrequired bool
	Isactive         bool
}

func convertMachineDB(inp postgres.Machine) Machine {
	var out Machine
	out.ID = inp.ID
	out.Name = inp.Name
	if inp.Isrepairrequired.Valid {
		out.Isrepairrequired = inp.Isrepairrequired.Bool
	}
	out.Isactive = inp.Isactive.Bool
	return out
}

type ShiftWorker struct {
	Shiftid int64
	Userid  int64
}

func convertShiftWorkerDB(inp postgres.ShiftWorker) ShiftWorker {
	var out ShiftWorker
	out.Shiftid = inp.Shiftid
	out.Userid = inp.Userid
	return out
}

type ShiftTask struct {
	Shiftid int64
	Taskid  int64
}

func convertShiftTaskDB(inp postgres.ShiftTask) ShiftTask {
	var out ShiftTask
	out.Shiftid = inp.Shiftid
	out.Taskid = inp.Taskid
	return out
}
