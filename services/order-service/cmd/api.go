/*
Copyright © 2024 Tejiri tejiriaustin123@gmail.com
*/
package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/tejiriaustin/literate-robot/core/config"
	"github.com/tejiriaustin/literate-robot/core/database"
	"google.golang.org/grpc"

	"order-service/controller"
	"order-service/repository"
	"order-service/service"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts Order Service API",
	Run:   startAPI,
}

func startAPI(cmd *cobra.Command, args []string) {
	cfg := setupEnvironment()

	dbCfg := &database.Config{
		Host:     cfg.GetAsString("HOST"),
		Port:     cfg.GetAsString("DB_PORT"),
		User:     cfg.GetAsString("USERNAME"),
		Password: cfg.GetAsString("PASSWORD"),
		DBName:   cfg.GetAsString("DB_NAME"),
	}
	dbConn, err := database.Initialize(dbCfg)
	if err != nil {
		return
	}

	repo := repository.NewOrderServiceRepository(dbConn)

	src := service.NewOrderService()

	ctlr := controller.NewOrderController()

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
		AddEnv("PORT", config.MustGetEnv("PORT")).
		AddEnv("PORT", config.MustGetEnv("HOST")).
		AddEnv("PORT", config.MustGetEnv("DB_PORT")).
		AddEnv("PORT", config.MustGetEnv("USERNAME")).
		AddEnv("PORT", config.MustGetEnv("PASSWORD")).
		AddEnv("PORT", config.MustGetEnv("DB_NAME"))

	return &cfg
}
