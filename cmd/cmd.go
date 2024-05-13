package cmd

import (
	"context"
	"os"

	"gateway-grpc/lib/pkg/server"

	"gateway-grpc/route"

	_ "gateway-grpc/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli"
)

func runCommand() func(*cli.Context) error {
	return func(c *cli.Context) error {

		srv := server.NewServer()
		ctx := context.Background()

		// route for user service
		route.SetupRouterUser(ctx, srv)

		//swagger api documentation
		if os.Getenv("ENV") != "production" {
			srv.GET("/swagger/*", echoSwagger.WrapHandler)
		}

		// start serve
		port := "PORT"
		portNumber := os.Getenv(port)
		srv.Logger.Fatal(srv.Start(":" + portNumber))

		return nil
	}
}

func Cli() *cli.App {
	clientApp := cli.NewApp()
	clientApp.Name = os.Getenv("SERVICE_NAME")
	clientApp.Version = os.Getenv("SERVICE_VERSION")
	clientApp.Action = runCommand()
	clientApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "mode",
			// Required: false,
			Value: "private",
		},
	}
	return clientApp
}
