package handler_output

type UserOutput struct {
	ID       int64
	Bitrixid int64
	Name     string
	Role     string
}

type ShiftOutput struct {
	ID            int64
	Machineid     int64
	ShiftMaster   int64
	Createdat     string
	Isactive      bool
	Deactivatedat string
}
type TaskOutput struct {
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

type MachineOutput struct {
	ID               int64
	Name             string
	Isrepairrequired bool
	Isactive         bool
}
