package quotesmanager

import (
	"bufio"
	"math/rand"
	"os"
)

type quotesManager struct {
	quotes []string
}

func New(filePath string) (*quotesManager, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &quotesManager{quotes: quotes}, nil
}

func (cm *quotesManager) GetQuote() string {
	if len(cm.quotes) == 0 {
		return ""
	}

	return cm.quotes[rand.Intn(len(cm.quotes))]
}
