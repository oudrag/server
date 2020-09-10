package main

import (
	"log"

	"github.com/oudrag/server/cmd/providers"
	"github.com/oudrag/server/internal/platform/app"
)

func main() {
	log.Fatal(
		app.NewIoC([]app.ServiceProvider{
			new(providers.DatabaseServiceProvider),
			new(providers.GraphServiceProvider),
			new(providers.CQRSServiceProvider),
			new(providers.RoutingServiceProvider),
		}).Boot(),
	)
}
