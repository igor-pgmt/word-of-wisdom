package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ListenAddr    string
	WowFile       string
	MaxWorkers    uint64
	ConnDeadline  time.Duration
	PowDifficulty uint8
}

func LoadConf() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		return nil, errors.New("LISTEN_ADDR is empty")
	}

	wowFile := os.Getenv("WOW_FILE")
	if wowFile == "" {
		return nil, errors.New("WOW_FILE is empty")
	}

	maxWorkersStr := os.Getenv("MAX_WORKERS")
	if maxWorkersStr == "" {
		return nil, errors.New("MAX_WORKERS is empty")
	}
	maxWorkers, err := strconv.ParseUint(maxWorkersStr, 10, 0)
	if err != nil {
		log.Fatalf("Invalid integer format for MAX_WORKERS: %v", err)
	}

	connDeadlineStr := os.Getenv("CONN_DEADLINE")
	if connDeadlineStr == "" {
		return nil, errors.New("CONN_DEADLINE is empty")
	}
	connDeadline, err := time.ParseDuration(connDeadlineStr)
	if err != nil {
		log.Fatalf("Invalid CONN_DEADLINE format for time.Duration: %v", err)
	}

	powDifficultyStr := os.Getenv("POW_DIFFICULTY")
	if powDifficultyStr == "" {
		return nil, errors.New("POW_DIFFICULTY is empty")
	}
	powDifficulty, err := strconv.ParseUint(powDifficultyStr, 10, 0)
	if err != nil {
		log.Fatalf("Invalid integer format for POW_DIFFICULTY: %v", err)
	}

	return &Config{
		ListenAddr:    listenAddr,
		WowFile:       wowFile,
		MaxWorkers:    maxWorkers,
		ConnDeadline:  connDeadline,
		PowDifficulty: uint8(powDifficulty),
	}, nil
}
