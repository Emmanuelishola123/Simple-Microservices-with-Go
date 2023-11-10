package main

import (
	"context"
	"fmt"

	"github.com/Emmanuelishola123/microservices-with-go-reddis/app"
)

func main() {
	app := app.NewApp()

	err := app.Start((context.TODO()))

	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}
