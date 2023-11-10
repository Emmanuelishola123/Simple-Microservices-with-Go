package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Emmanuelishola123/microservices-with-go-reddis/app"
)

func main() {
	app := app.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start((ctx))

	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}
