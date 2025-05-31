package convertor

import "github.com/jackc/pgx/v5/pgtype"

func PGInt64(x *int64) pgtype.Int8 {
	var pgX pgtype.Int8
	if x == nil {
		pgX.Valid = false
	} else {
		pgX.Valid = true
		pgX.Int64 = *x
	}
	return pgX
}

func PGBool(x *bool) pgtype.Bool {
	var pgX pgtype.Bool
	if x == nil {
		pgX.Valid = false
	} else {
		pgX.Valid = true
		pgX.Bool = *x
	}
	return pgX
}
