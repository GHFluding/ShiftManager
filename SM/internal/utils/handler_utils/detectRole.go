package handler_utils

import "sm/internal/database/postgres"

func DetectUserRole(sRole string) (postgres.Userrole, bool) {
	switch sRole {
	case "engineer":
		return postgres.UserroleEngineer, true
	case "worker":
		return postgres.UserroleWorker, true
	case "master":
		return postgres.UserroleMaster, true
	case "manager":
		return postgres.UserroleManager, true
	case "admin":
		return postgres.UserroleAdmin, true
	default:
		return postgres.Userrole(""), false
	}
}
