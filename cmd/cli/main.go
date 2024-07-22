package main

import (
	"bufio"
	"fmt"
	"go.vardan.dev/perdis/internal/database"
	"log/slog"
	"os"
)

const (
	logFilePath = "perdis.log"
)

func main() {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open log file: %v\n", err)
		return
	}

	slog.SetLogLoggerLevel(slog.LevelDebug)
	logger := slog.New(slog.NewTextHandler(logFile, nil))

	db, err := database.Start(logger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start database: %v\n", err)
		return
	}

	fmt.Println("perdis interactive console")

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read input: %v\n", err)
			return
		}
		fmt.Println(db.Execute(line))
	}
}
