package main

import (
	"bytes"
	"context"
	"encoding/json"
	"html/template"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const tplData = `{{ range . }}
<div class="check">
<div class="state state-{{ .StateClass }}">{{ .State }}</div>
<div class="service">{{ .Service }}</div>
<div class="uptime">{{ .Uptime | printf "%.03f"}} %</div>
</div>
{{ end }}
`

var tpl = template.Must(template.New("index.html").Parse(tplData))

var (
	apiKey        = os.Getenv("UPDOWN_APIKEY")
	domain        = os.Getenv("UPDOWN_DOMAIN")
	cacheDuration = 5 * time.Minute
)

var (
	cache     string
	cacheTime time.Time
	cacheMut  sync.RWMutex
)

func main() {
	// if cs, err := currentStatus(apiKey, domain); err == nil {
	// 	tpl.Execute(os.Stdout, cs)
	// }
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

func currentStatus(apiKey, domain string) (string, error) {
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
		return "", err
	}

	checks = filterPublicDomain(checks, domain)
	status := sortedStatus(checks)
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, status); err != nil {
		return "", err
	}

	cache = buf.String()
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
