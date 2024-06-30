package service

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type ProofFinder interface {
	FindProof(challenge string) string
}

type service struct {
	address string
	pf      ProofFinder
}

func New(adress string, pf ProofFinder) *service {
	return &service{address: adress, pf: pf}
}

func (s *service) Run(ctx context.Context) {
	log := logrus.New()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var conn net.Conn
	var err error
	for {
		select {
		case <-ctx.Done():
			log.Info("Shutting down client...")
			return

		case <-ticker.C:
			conn, err = net.Dial("tcp", s.address)
			if err != nil {
				log.Errorf("Failed to connect to server: %v", err)
				continue
			}
			scanner := bufio.NewScanner(conn)

			if scanner.Scan() {
				challengeMsg := scanner.Text()
				challenge := strings.TrimPrefix(challengeMsg, "Challenge: ")

				proof := s.pf.FindProof(challenge)
				fmt.Fprintf(conn, "%s\n", proof)

				if scanner.Scan() {
					log.Info(scanner.Text())
				}
			}

			conn.Close()
		}
	}
}
