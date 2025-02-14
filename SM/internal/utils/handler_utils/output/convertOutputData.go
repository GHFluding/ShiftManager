package handler_output

import (
	"errors"
	"sm/internal/database/postgres"
)

func ConvertToOut(dest interface{}) (interface{}, error) {
	switch inp := dest.(type) {
	case postgres.User:
		{
			var out UserOutput
			out.ID = inp.ID
			out.Bitrixid = inp.ID
			out.Name = inp.Name
			out.Role = string(inp.Role)
			return out, nil
		}
	case postgres.Shift:
		{
			var out ShiftOutput
			out.ID = inp.ID
			out.Machineid = inp.Machineid
			out.ShiftMaster = inp.ShiftMaster
			out.Isactive = inp.Isactive.Bool
			out.Createdat = inp.Createdat.Time.String()
			out.Deactivatedat = inp.Createdat.Time.String()
			return out, nil
		}
	case postgres.ShiftTask:
		return inp, nil
	case postgres.ShiftWorker:
		return inp, nil
	case postgres.Task:
		{
			var out TaskOutput
			out.ID = inp.ID
			out.Machineid = inp.Machineid
			if !inp.Shiftid.Valid {
				//what to return if there is no data
				out.Shiftid = 0
			}
			out.Shiftid = inp.Shiftid.Int64
			out.Frequency = string(inp.Frequency)
			out.Taskpriority = string(inp.Taskpriority)
			out.Description = inp.Description
			out.Createdby = inp.Createdby
			out.Createdat = inp.Createdat.Time.String()
			if !inp.Verifiedby.Valid {
				//what to return if there is no data
				out.Verifiedby = 0
			}
			out.Verifiedby = inp.Verifiedby.Int64
			out.Verifiedat = inp.Verifiedat.Time.String()
			if !inp.Verifiedby.Valid {
				//what to return if there is no data
				out.Completedby = 0
			}
			out.Completedby = inp.Completedby.Int64
			out.Completedat = inp.Completedat.Time.String()
			out.Status = string(inp.Status)
			out.Comment = inp.Comment.String
			if !inp.Verifiedby.Valid {
				//what to return if there is no data
				out.Movedinprogressby = 0
			}
			out.Movedinprogressby = inp.Movedinprogressby.Int64
			out.Movedinprogressat = inp.Movedinprogressat.Time.String()
			return out, nil
		}
	case postgres.Machine:
		{
			var out MachineOutput
			out.ID = inp.ID
			out.Isactive = inp.Isactive.Bool
			out.Isrepairrequired = inp.Isrepairrequired.Bool
			out.Name = inp.Name
			return out, nil
		}
	default:
		return nil, errors.New("invalid input to convert output")

	}
}

func ListConvert(dest []interface{}) ([]interface{}, error) {
	var usersOut []interface{}
	inpSlice := dest
	switch inp := dest[0].(type) {
	case postgres.User:
		{
			for i := range inp {
				userOut, err := ConvertToOut(inp[i])
				usersOut = append(usersOut, userOut)
				if err != nil {
					return nil, err
				}
			}
			return usersOut, nil
		}
	default:
		{
			for i := range inp {
				userOut, err := ConvertToOut(inp[i])
				usersOut = append(usersOut, userOut)
				if err != nil {
					return nil, err
				}
			}
			return usersOut, nil
		}
	}
}
