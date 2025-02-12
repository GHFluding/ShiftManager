// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgres

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Taskfrequency string

const (
	TaskfrequencyOneTime   Taskfrequency = "one_time"
	TaskfrequencyDaily     Taskfrequency = "daily"
	TaskfrequencyWeekly    Taskfrequency = "weekly"
	TaskfrequencyMonthly   Taskfrequency = "monthly"
	TaskfrequencyQuarterly Taskfrequency = "quarterly"
)

func (e *Taskfrequency) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Taskfrequency(s)
	case string:
		*e = Taskfrequency(s)
	default:
		return fmt.Errorf("unsupported scan type for Taskfrequency: %T", src)
	}
	return nil
}

type NullTaskfrequency struct {
	Taskfrequency Taskfrequency
	Valid         bool // Valid is true if Taskfrequency is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTaskfrequency) Scan(value interface{}) error {
	if value == nil {
		ns.Taskfrequency, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Taskfrequency.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTaskfrequency) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Taskfrequency), nil
}

type Taskpriority string

const (
	TaskpriorityLow     Taskpriority = "low"
	TaskpriorityMiddle  Taskpriority = "middle"
	TaskpriorityHotTask Taskpriority = "hot_task"
	TaskpriorityHotFix  Taskpriority = "hot_fix"
)

func (e *Taskpriority) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Taskpriority(s)
	case string:
		*e = Taskpriority(s)
	default:
		return fmt.Errorf("unsupported scan type for Taskpriority: %T", src)
	}
	return nil
}

type NullTaskpriority struct {
	Taskpriority Taskpriority
	Valid        bool // Valid is true if Taskpriority is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTaskpriority) Scan(value interface{}) error {
	if value == nil {
		ns.Taskpriority, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Taskpriority.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTaskpriority) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Taskpriority), nil
}

type Taskstatus string

const (
	TaskstatusTodo       Taskstatus = "todo"
	TaskstatusInProgress Taskstatus = "inProgress"
	TaskstatusFailed     Taskstatus = "failed"
	TaskstatusCompleted  Taskstatus = "completed"
	TaskstatusVerified   Taskstatus = "verified"
)

func (e *Taskstatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Taskstatus(s)
	case string:
		*e = Taskstatus(s)
	default:
		return fmt.Errorf("unsupported scan type for Taskstatus: %T", src)
	}
	return nil
}

type NullTaskstatus struct {
	Taskstatus Taskstatus
	Valid      bool // Valid is true if Taskstatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTaskstatus) Scan(value interface{}) error {
	if value == nil {
		ns.Taskstatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Taskstatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTaskstatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Taskstatus), nil
}

type Userrole string

const (
	UserroleEngineer Userrole = "engineer"
	UserroleWorker   Userrole = "worker"
	UserroleMaster   Userrole = "master"
	UserroleManager  Userrole = "manager"
	UserroleAdmin    Userrole = "admin"
)

func (e *Userrole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Userrole(s)
	case string:
		*e = Userrole(s)
	default:
		return fmt.Errorf("unsupported scan type for Userrole: %T", src)
	}
	return nil
}

type NullUserrole struct {
	Userrole Userrole
	Valid    bool // Valid is true if Userrole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserrole) Scan(value interface{}) error {
	if value == nil {
		ns.Userrole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Userrole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserrole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Userrole), nil
}

type Machine struct {
	ID               int64
	Name             string
	Isrepairrequired pgtype.Bool
	Isactive         pgtype.Bool
}

type Shift struct {
	ID            int64
	Machineid     int64
	ShiftMaster   int64
	Createdat     pgtype.Date
	Isactive      pgtype.Bool
	Deactivatedat pgtype.Date
}

type ShiftTask struct {
	Shiftid int64
	Taskid  int64
}

type ShiftWorker struct {
	Shiftid int64
	Userid  int64
}

type Task struct {
	ID                int64
	Machineid         int64
	Shiftid           pgtype.Int8
	Frequency         Taskfrequency
	Taskpriority      Taskpriority
	Description       string
	Createdby         int64
	Createdat         pgtype.Date
	Verifiedby        pgtype.Int8
	Verifiedat        pgtype.Date
	Completedby       pgtype.Int8
	Completedat       pgtype.Date
	Status            Taskstatus
	Comment           pgtype.Text
	Movedinprogressby pgtype.Int8
	Movedinprogressat pgtype.Date
}

type User struct {
	ID       int64
	Bitrixid int64
	Name     string
	Role     Userrole
}
