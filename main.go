package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jcobhams/mmhmm/config"
	"github.com/jcobhams/mmhmm/database"
	"github.com/jcobhams/mmhmm/env"
	"github.com/jcobhams/mmhmm/logging"
	"github.com/jcobhams/mmhmm/repositories"
	"github.com/jcobhams/mmhmm/server"
	"github.com/jcobhams/mmhmm/services"
)

func main() {
	logging.InitializeFromEnv()
	ctx, _ := logging.Logger(context.Background(), "main")

	conf := config.New(env.ConfigProvider(ctx))

	db := database.New()

	serviceContainer := services.NewContainer()
	repoContainer := repositories.NewContainer(db)

	//Start go routine ticker to print contents of db -- Debug Purpose
	go func() {
		for {
			fmt.Println(db.All())
			<-time.After(time.Second * 5)
		}
	}()

	server.Start(ctx, serviceContainer, repoContainer, conf)
}
