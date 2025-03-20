package config

type APIRoutes struct {
	Shift       string
	Machine     string
	Task        string
	ShiftWorker string
}

var Routes = APIRoutes{
	Shift:       "/shift",
	Machine:     "/machine",
	Task:        "/task",
	ShiftWorker: "/shift/worker",
}
