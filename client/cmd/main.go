package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/igor-pgmt/word-of-wisdom/client/internal/service"
	"github.com/igor-pgmt/word-of-wisdom/pkg/pow"
)

func main() {
	log := logrus.New()

	conf, err := LoadConf()
	if err != nil {
		log.Fatalf("Failed to load envs: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Info("Received shutdown signal")
		cancel()
	}()

	pow := pow.New(conf.PowDifficulty)

	srv := service.New(conf.SrvAddr, pow)
	srv.Run(ctx)
}
