package main

import (
	"context"
	"github.com/tclutin/classconnect-api/internal/app"
)

//	@title			classconnect-api
//	@version		1.0
//	@description	API Server for the ClassConnect application

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.apiKey	Bearer
// @in							header
// @name						Authorization
// @description				Type "Bearer + jwtToken"
func main() {
	app.New().Run(context.Background())
}
