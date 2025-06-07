package grpchandler

import (
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/app"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

func (grpcS *gRPCServer) RunMachine() {
	application := app.New(grpcS.log, grpcS.port, grpcS.db, models.MachineServer)
	application.Run(grpcS.db)
}
