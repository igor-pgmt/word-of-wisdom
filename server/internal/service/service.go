package service

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Quoter interface {
	GetQuote() string
}

type ProofOfWorker interface {
	GenerateChallenge() string
	ValidateProof(challenge, proof string) bool
}

type service struct {
	address    string
	deadline   time.Duration
	workerPool chan struct{}
	cm         Quoter
	pow        ProofOfWorker
}

func New(adress string, maxWorkers uint64, deadline time.Duration, cm Quoter, pow ProofOfWorker) *service {
	return &service{
		address:    adress,
		deadline:   deadline,
		workerPool: make(chan struct{}, maxWorkers),
		cm:         cm,
		pow:        pow,
	}
}

func (s *service) Run(ctx context.Context, log *logrus.Logger) {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Errorf("Failed to listen on %s: %v", s.address, err)
	}
	defer listener.Close()

	log.Infof("Server listening on %s", s.address)

	wg := &sync.WaitGroup{}

loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Shutting down server...")
			break loop

		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Errorf("Failed to accept connection: %v", err)
				continue
			}
			wg.Add(1)

			s.workerPool <- struct{}{}
			go s.handleConnection(conn, wg)
		}
	}

	wg.Wait()

	log.Info("Server stopped")
}

func (s *service) handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()

	challenge := s.pow.GenerateChallenge()
	fmt.Fprintf(conn, "Challenge: %s\n", challenge)

	conn.SetReadDeadline(
		time.Now().Add(s.deadline),
	)

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		proof := scanner.Text()
		if s.pow.ValidateProof(challenge, proof) {
			fmt.Fprintf(conn, "Word of Wisdom: %s\n", s.cm.GetQuote())
		} else {
			fmt.Fprintln(conn, "Invalid proof of work.")
		}
	} else {
		fmt.Fprintln(conn, "Timeout or connection error.")
	}

	<-s.workerPool
}
