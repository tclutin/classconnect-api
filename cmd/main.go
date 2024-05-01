package main

import (
	_ "classconnect-api/docs"
	"classconnect-api/internal/app"
	"context"
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
