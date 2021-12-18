package main

import (
	"flag"
	"fmt"
	"github-program-rate/pkg/config"
	"github-program-rate/pkg/domain/application"
	"github-program-rate/pkg/gateway"
	logger2 "github-program-rate/pkg/logger"
	"github.com/rs/zerolog"
	"io/ioutil"
	"log"
	"os"
)

var (
	version = "0.1.0"
	logger  *zerolog.Logger
)

func init() {

	file, err := os.Open("rate.conf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	config := config.NewConfig(buffer)
	fmt.Println(config)

	logger, err = logger2.NewLogger(config)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info().Str("Title", config.Title).Msg("Config")
	logger.Info().Str("LogLevel", config.LogInfo.Level).Msg("Config")
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Printf("g-sig version is %s", version)
		return
	}

	// UseCase
	drawSVGUseCase := application.NewDrawSVGUseCase(logger)

	// Controller
	userHistoryHandler := gateway.NewUserHistoryHandler(drawSVGUseCase, logger)

	logger.Info().Str("Addr", os.Getenv("PORT")).Msg("Serve is running")

	server := gateway.NewServer(logger, userHistoryHandler)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal().Err(err).Msg("ListenAndServe:Error")
	}
}
