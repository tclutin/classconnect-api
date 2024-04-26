package main

import (
	"classconnect-api/internal/app"
	"context"
)

func main() {
	app.New().Run(context.Background())
}
