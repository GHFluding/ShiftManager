package config

type Routing struct {
	baseURL string
}

func (r *Routing) GetBaseURL() string {
	return r.baseURL
}

func (r *Routing) GetTaskBaseURL() string {
	return (r.baseURL + "/task")
}

func (r *Routing) GetShiftBaseURL() string {
	return (r.baseURL + "/shifts")
}

func (r *Routing) GetMachineBaseURL() string {
	return (r.baseURL + "/machine")
}

func (r *Routing) GetUserBaseURL() string {
	return (r.baseURL + "/users")
}

func (r *Routing) GetShiftTaskBaseURL() string {
	return (r.baseURL + "/shifts/tasks")
}

func (r *Routing) GetShiftWorkerBaseURL() string {
	return (r.baseURL + "/shifts/workers")
}
