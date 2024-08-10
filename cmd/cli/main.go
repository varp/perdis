package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"go.vardan.dev/perdis/internal/config"
	"go.vardan.dev/perdis/internal/database"
	"log"
	"log/slog"
	"os"
)

const (
	logFilePath = "perdis.log"

	configEnvPrefix = "PERDIS"
	configName      = "perdis"
	configExt       = "yaml"
)

func main() {

	viper.SetConfigName(configName)
	viper.SetConfigType(configExt)
	viper.AddConfigPath(".")
	viper.SetEnvPrefix(configEnvPrefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("failed to read configuration %s", err)
		return
	}

	cfg := config.Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Printf("failed to parase configuration %s", err)
		return
	}

	fmt.Println(cfg)
	return

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v\n", err)
		return
	}

	slog.SetLogLoggerLevel(slog.LevelDebug)
	logger := slog.New(slog.NewTextHandler(logFile, nil))

	db, err := database.Start(logger)
	if err != nil {
		log.Fatalf("failed to start database: %v\n", err)
		return
	}

	fmt.Println("perdis interactive console")

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read input: %v\n", err)
			return
		}
		fmt.Println(db.Execute(line))
	}
}
