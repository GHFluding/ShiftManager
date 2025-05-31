package grpchandler

import (
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/app"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

func (grpcS *gRPCServer) RunTask() {
	application := app.New(grpcS.log, grpcS.port, grpcS.db, models.TaskServer)
	application.Run(grpcS.db)
}
