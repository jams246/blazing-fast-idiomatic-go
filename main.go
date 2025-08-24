package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type SearchResponse struct {
	TotalCount int `json:"total_count"`
}

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	query := url.QueryEscape(`"blazing fast" OR "idiomatic" language:go in:readme in:description`)
	apiURL := fmt.Sprintf("https://api.github.com/search/repositories?q=%s", query)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		log.Fatalf("Failed to create idiomatic request: %v", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "blazing-fast-counter/1.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request wasn't blazing fast enough: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GitHub said no (status %d). Maybe we're too fast?", resp.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20)) // 1MB is plenty
	if err != nil {
		log.Fatalf("Failed to read idiomatically: %v", err)
	}

	var result SearchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("JSON parsing failed (not idiomatic enough?): %v", err)
	}

	fmt.Printf("🔥 Blazingly Fast™ and Idiomatic™ Go repos: %d\n", result.TotalCount)
}
