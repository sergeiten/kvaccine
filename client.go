package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type Stop struct {
	error
}

func NewTimeoutClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

func Retry(attempts int, sleep time.Duration, f func() (*http.Response, error)) (*http.Response, error) {
	resp, err := f()
	if err != nil {
		if s, ok := err.(*Stop); ok {
			// Return the original error for later checking
			return nil, s.error
		}

		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)

			return Retry(attempts, 2*sleep, f)
		}
		return nil, err
	}

	return resp, nil
}

func RetryRequest(client *http.Client, request *http.Request, attempts int, sleep time.Duration) (*http.Response, error) {
	return Retry(attempts, sleep, func() (*http.Response, error) {
		resp, err := client.Do(request)
		if err != nil {
			return nil, err
		}

		s := resp.StatusCode
		switch {
		case s >= 500:
			// Retry
			return nil, fmt.Errorf("server error: %v", s)
		case s >= 400:
			// Don't retry, it was client's fault
			return nil, Stop{fmt.Errorf("client error: %v", s)}
		default:
			// Happy
			return resp, nil
		}
	})
}
