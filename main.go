package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const API_URL = "https://vaccine-map.kakao.com/api/v2/vaccine/left_count_by_coords"

type response struct {
	Organizations []struct {
		OrgCode    string  `json:"orgCode"`
		OrgName    string  `json:"orgName"`
		Address    string  `json:"address"`
		X          float64 `json:"x"`
		Y          float64 `json:"y"`
		Status     string  `json:"status"`
		LeftCounts int     `json:"leftCounts"`
	} `json:"organizations"`
}

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	tlg := NewTelegramLogger(cfg.TelegramToken, cfg.TelegramChatID)

	c := NewTimeoutClient()
	locations := map[string][]byte{
		"seoul": []byte(`
{
  "bottomRight": {
    "x": 127.225871,
    "y": 37.385768
  },
  "onlyLeft": true,
  "order": "latitude",
  "topLeft": {
    "x": 126.630329,
    "y": 37.672568
  }
}`),
	}

	timePassed := 0
	sleep := 1
	sendMessageAfter := 60

	for {
		for location, data := range locations {
			req, err := newRequest(API_URL, data)
			if err != nil {
				log.Fatalf("failed to create request: %v", err)
			}

			resp, err := c.Do(req)
			if err != nil {
				log.Fatalf("failed to do request: %v", err)
			}

			defer resp.Body.Close()

			dat, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("failed to read response body: %v", err)
			}

			var respJson response
			err = json.Unmarshal(dat, &respJson)
			if err != nil {
				log.Fatalf("failed to unmarshal json: %v", err)
			}

			if len(respJson.Organizations) == 0 {
				log.Printf("There is no vaccine left for %s location\n", location)
				continue
			}

			for _, org := range respJson.Organizations {
				tlg.Logf("%s", org.Address)
				tlg.Logf("Hospital: %s\nAddress: %s\nHas: %d vaccine left", org.OrgName, org.Address, org.LeftCounts)
			}
		}

		timePassed += sleep
		log.Printf("Sleep for %d minutes\n", sleep)
		time.Sleep(time.Duration(sleep) * time.Minute)

		if timePassed >= sendMessageAfter {
			tlg.Logf("There is still no vaccine")
			timePassed = 0
		}
	}
}

func newRequest(url string, data []byte) (*http.Request, error) {
	buf := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Length", strconv.Itoa(buf.Len()))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Origin", "https://vaccine-map.kakao.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.1 EVENT_LINKTAB")
	req.Header.Set("Referer", "https://vaccine-map.kakao.com/map2?v=1")

	return req, nil
}
