package main

import (
	"context"
	"encoding/json"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	apiKey        = os.Getenv("UPDOWN_APIKEY")
	domain        = os.Getenv("UPDOWN_DOMAIN")
	cacheDuration = 5 * time.Minute
)

var (
	cache     []status
	cacheTime time.Time
	cacheMut  sync.RWMutex
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	ss, err := currentStatus(apiKey, domain)
	if err != nil {
		return nil, err
	}
	bs, _ := json.MarshalIndent(ss, "", " ")
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		Body: string(bs),
	}, nil
}

func currentStatus(apiKey, domain string) ([]status, error) {
	cacheMut.RLock()
	if time.Since(cacheTime) < cacheDuration {
		cacheMut.RUnlock()
		return cache, nil
	}
	cacheMut.RUnlock()

	cacheMut.Lock()
	defer cacheMut.Unlock()
	if time.Since(cacheTime) < cacheDuration {
		return cache, nil
	}

	checks, err := getChecks(apiKey)
	if err != nil {
		return nil, err
	}

	checks = filterPublicDomain(checks, domain)
	cache = sortedStatus(checks)
	cacheTime = time.Now()
	return cache, nil
}

func sortedStatus(checks []check) []status {
	ss := make([]status, len(checks))
	for i, c := range checks {
		ss[i] = c.status()
	}
	sort.Slice(ss, func(a, b int) bool {
		if ss[a].OK != ss[b].OK {
			return !ss[a].OK
		}
		return ss[a].Service < ss[b].Service
	})
	return ss
}
