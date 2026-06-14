package app

import (
	"fmt"
	"log"
	"os"

	"github.com/Junaidmdv/goalcircle-tournament_service/internal/config"
	"github.com/Junaidmdv/goalcircle-tournament_service/internal/infrastructure/server"
	"github.com/Junaidmdv/goalcircle-tournament_service/pkg/logger"
)

func main() {

	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("failed to initialise logger: %v", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "logger sync error: %v\n", err)
		}
	}()

	cnfg, errs := config.LoadConfig().
		WithGrpcServer().
		Build()

	if errs != nil {
		for _, err := range errs {
			logger.Error("configration error", "error", err)
			fmt.Println()
		}
		return
	}

	server := server.NewGRPCServer(cnfg)

}
