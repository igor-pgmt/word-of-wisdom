package main

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SrvAddr       string
	PowDifficulty uint8
}

func LoadConf() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	srvAddr := os.Getenv("SRV_ADDR")
	if srvAddr == "" {
		return nil, errors.New("SRV_ADDR is empty")
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
		SrvAddr:       srvAddr,
		PowDifficulty: uint8(powDifficulty),
	}, nil
}
