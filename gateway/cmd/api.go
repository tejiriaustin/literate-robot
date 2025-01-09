// Copyright © 2024 Tejiri tejiriaustin123@gmail.com

package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/tejiriaustin/literate-robot/core/config"
	"google.golang.org/grpc"

	"gateway/controller"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the Gateway",
	Run:   startAPI,
}

func startAPI(cmd *cobra.Command, args []string) {
	cfg := setupEnvironment()

	ctlr := controller.NewGatewayController()

	listener, err := net.Listen("tcp", cfg.GetAsString("PORT"))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", cfg.GetAsString("PORT"), err)
	}
	grpcServer := grpc.NewServer()

	gateway.RegisterGatewayServiceServer(grpcServer, ctlr)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func setupEnvironment() *config.Environment {
	cfg := config.NewEnvironment().
		AddEnv("PORT", config.MustGetEnv("PORT"))

	return &cfg
}
