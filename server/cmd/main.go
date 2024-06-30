package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/igor-pgmt/word-of-wisdom/pkg/pow"
	"github.com/igor-pgmt/word-of-wisdom/pkg/quotesmanager"

	"github.com/igor-pgmt/word-of-wisdom/server/internal/service"
)

func main() {
	log := logrus.New()

	conf, err := LoadConf()
	if err != nil {
		log.Fatalf("Failed to load envs: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	qm, err := quotesmanager.New(conf.WowFile)
	if err != nil {
		log.Fatalf("Failed to create CitateManager: %v", err)
	}

	pow := pow.New(conf.PowDifficulty)
	if err != nil {
		log.Fatalf("Failed to create CitateManager: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Info("Received shutdown signal")
		cancel()
	}()

	srv := service.New(conf.ListenAddr, conf.MaxWorkers, conf.ConnDeadline, qm, pow)
	srv.Run(ctx, log)
}
