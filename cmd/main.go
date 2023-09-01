package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/api"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/config"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/database"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/handler"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/repository"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/service"

	_ "github.com/MyLi2tlePony/AvitoInternshipGolang2023/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var configPath string

func init() {
	defaultConfigPath := path.Join("config", "config.toml")
	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to configuration file")
}

func main() {
	dbConfig, err := config.NewDatabaseConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()

	dbConnectionString := dbConfig.GetConnectionString()
	postgresDsn := os.Getenv("POSTGRES_DSN")
	if postgresDsn != "" {
		dbConnectionString = postgresDsn
	}

	pg, err := database.NewPG(ctx, dbConnectionString)
	if err != nil {
		fmt.Println(err)
		return
	}

	segmentRepo := repository.NewSegmentRepository(ctx, pg)
	userRepo := repository.NewUserRepository(ctx, pg)

	segmentService := service.NewSegmentService(segmentRepo)
	userService := service.NewUserService(userRepo)

	segmentHandler := handler.NewSegmentHandler(segmentService)
	userHandler := handler.NewUserHandler(userService)

	srv := echo.New()
	srv.GET("/swagger/*", echoSwagger.WrapHandler)

	api.SetupSegmentRoutes(srv, segmentHandler)
	api.SetupUserRoutes(srv, userHandler)

	go func() {
		for {
			if err := userRepo.DeleteOldSegments(); err != nil {
				fmt.Println(err)
			}

			time.Sleep(60 * time.Second)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("failed to stop serv: " + err.Error())
		}
	}()

	logConfig, err := config.NewLoggerConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	srv.Logger.SetLevel(logConfig.GetLevel())

	srvConfig, err := config.NewServerConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = srv.Start(srvConfig.GetHostPort())
	if err != nil {
		fmt.Println(err)
	}
}
