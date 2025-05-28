package grpchandler

import (
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/app"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

func (grpcS *gRPCServer) RunShift() {
	application := app.New(grpcS.log, grpcS.port, grpcS.db, models.TaskServer)
	application.Run(grpcS.db)
}
