package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/tejiriaustin/literate-robot/core/config"
	"github.com/tejiriaustin/literate-robot/core/database"
	"github.com/tejiriaustin/literate-robot/core/logger"
	"google.golang.org/grpc"

	"user-service/controller"
	"user-service/enviroment"
	"user-service/repository"
	"user-service/service"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts User Service API",
	Run:   startAPI,
}

func startAPI(cmd *cobra.Command, args []string) {
	cfg := setupEnvironment()

	nlog := logger.NewZeroLogger()

	dbCfg := &database.Config{
		Host:     cfg.GetAsString("HOST"),
		Port:     cfg.GetAsString("DB_PORT"),
		User:     cfg.GetAsString("USERNAME"),
		Password: cfg.GetAsString("PASSWORD"),
		DBName:   cfg.GetAsString("DB_NAME"),
	}
	dbConn, err := database.Initialize(dbCfg)
	if err != nil {
		nlog.Fatal("Failed to connect to database", logger.WithField("err", err))
		return
	}

	userRepo := repository.NewUserServiceRepository(dbConn)

	userService := service.NewUserService()

	controller := controller.NewUserController(userService, userRepo)

	lis, err := net.Listen("tcp", cfg.GetAsString("PORT"))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", cfg.GetAsString("PORT"), err)
	}
	grpcServer := grpc.NewServer()
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func setupEnvironment() *config.Environment {
	cfg := config.NewEnvironment().
		AddEnv(enviroment.Port, config.MustGetEnv(enviroment.Port)).
		AddEnv(enviroment.Host, config.MustGetEnv(enviroment.Host)).
		AddEnv(enviroment.DatabaseName, config.MustGetEnv(enviroment.DatabaseName)).
		AddEnv(enviroment.DatabaseUsername, config.MustGetEnv(enviroment.DatabaseUsername)).
		AddEnv(enviroment.DatabasePassword, config.MustGetEnv(enviroment.DatabasePassword))

	return &cfg
}
