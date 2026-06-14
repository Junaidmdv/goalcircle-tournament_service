package app

import (
	"fmt"

	"github.com/Junaidmdv/goalcircle-tournament_service/internal/config"
	"github.com/Junaidmdv/goalcircle-tournament_service/internal/infrastructure/server"
)

func main() {  

	

	cnfg, errs := config.LoadConfig().
		WithGrpcServer().
		Build()

	if errs != nil {
		for _, err := range errs {
			//logger.Error("configration error", "error", err)
			fmt.Println()
		}
		return
	}

	cnfg

	server := server.NewGRPCServer()

}
