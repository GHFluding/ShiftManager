package models

import "context"

type Machine struct {
	Name             string
	IsRepairRequired *bool
	IsActive         *bool
}

type MachineSaver interface {
	SaveMachine(
		ctx context.Context,
		name string,
		isRepairRequired *bool,
		isActive *bool,
	) (
		id int64,
		err error,
	)
}
type MachineProvider interface {
	GetMachine(ctx context.Context, id int64) (Machine, error)
	// TODO: IsRepairRequired and isActive
}

type MachineDB struct {
	Provider MachineProvider
	Saver    MachineSaver
}
