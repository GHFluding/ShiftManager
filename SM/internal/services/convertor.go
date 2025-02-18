// convertOutputData.go
package services

import (
	"errors"
	"sm/internal/database/postgres"
)

// convert data type to output. For Example pgtype.Int8 to int64
func convertDataToTransport[T InputTypes, O OutputTypes](dest T) (O, error) {
	switch inp := any(dest).(type) {
	case postgres.User:
		var out UserDTO
		out.ID = inp.ID
		out.Bitrixid = inp.ID
		out.Name = inp.Name
		out.Role = string(inp.Role)
		return any(out).(O), nil
	case postgres.Shift:
		var out ShiftDTO
		out.ID = inp.ID
		out.Machineid = inp.Machineid
		out.ShiftMaster = inp.ShiftMaster
		out.Isactive = inp.Isactive.Bool
		out.Createdat = inp.Createdat.Time.String()
		out.Deactivatedat = inp.Deactivatedat.Time.String()
		return any(out).(O), nil
	case postgres.ShiftTask:
		var out ShiftTaskDTO
		out.Shiftid = inp.Shiftid
		out.Taskid = inp.Taskid
		return any(out).(O), nil
	case postgres.ShiftWorker:
		var out ShiftWorkerDTO
		out.Shiftid = inp.Shiftid
		out.Userid = inp.Userid
		return any(out).(O), nil
	case postgres.Task:
		var out TaskDTO
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
		out.Verifiedat = inp.Verifiedat.Time.String()
		if inp.Completedby.Valid {
			out.Completedby = inp.Completedby.Int64
		}
		out.Completedat = inp.Completedat.Time.String()
		out.Status = string(inp.Status)
		out.Comment = inp.Comment.String
		if inp.Movedinprogressby.Valid {
			out.Movedinprogressby = inp.Movedinprogressby.Int64
		}
		out.Movedinprogressat = inp.Movedinprogressat.Time.String()
		return any(out).(O), nil
	case postgres.Machine:
		var out MachineDTO
		out.ID = inp.ID
		out.Isactive = inp.Isactive.Bool
		out.Isrepairrequired = inp.Isrepairrequired.Bool
		out.Name = inp.Name
		return any(out).(O), nil
	default:
		return *new(O), errors.New("invalid input to convert output")
	}
}

// This function convert slice items to standard type
func convertListToTransport[T InputTypes, O OutputTypes](input []T) ([]O, error) {
	var sliceOut []O
	for _, item := range input {
		dataOut, err := convertDataToTransport[T, O](item)
		if err != nil {
			return nil, err
		}
		sliceOut = append(sliceOut, dataOut)
	}
	return sliceOut, nil
}
