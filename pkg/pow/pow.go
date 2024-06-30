package pow

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type proofOfWork struct {
	difficulty uint8
}

func New(difficulty uint8) *proofOfWork {
	return &proofOfWork{difficulty: difficulty}
}

func (pow *proofOfWork) GenerateChallenge() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (pow *proofOfWork) ValidateProof(challenge, proof string) bool {
	hash := sha256.Sum256([]byte(challenge + proof))
	hashStr := fmt.Sprintf("%x", hash)
	return strings.HasPrefix(hashStr, strings.Repeat("0", int(pow.difficulty)))
}

func (pow *proofOfWork) FindProof(challenge string) string {
	var proof int
	for {
		proofStr := strconv.Itoa(proof)
		if pow.ValidateProof(challenge, proofStr) {
			return proofStr
		}

		proof++
	}
}
