package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const LEFT_COUNT_URL = "https://vaccine-map.kakao.com/api/v2/vaccine/left_count_by_coords"
const ME_URL = "https://vaccine.kakao.com/api/v1/user"
const ORG_URL = "https://vaccine.kakao.com/api/v2/org/org_code"
const AGREEMENT_URL = "https://vaccine.kakao.com/api/v1/agreement"
const RESERVATION_URL = "https://vaccine.kakao.com/api/v1/reservation"

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

type userResponse struct {
	User struct {
		Name               string `json:"name"`
		Status             string `json:"status"`
		StatusDate         string `json:"statusDate"`
		StatusOrganization string `json:"statusOrganization"`
	} `json:"user"`
	Alarms       []interface{} `json:"alarms"`
	Reservations []interface{} `json:"reservations"`
}

type orgResponse struct {
	Organization struct {
		OrgCode     string  `json:"orgCode"`
		OrgName     string  `json:"orgName"`
		ConfirmId   string  `json:"confirmId"`
		PhoneNumber string  `json:"phoneNumber"`
		Address     string  `json:"address"`
		X           float64 `json:"x"`
		Y           float64 `json:"y"`
		OpenHour    struct {
			Date      string `json:"date"`
			DayOfWeek string `json:"dayOfWeek"`
			DayOff    bool   `json:"dayOff"`
			OpenHour  struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"openHour"`
			Lunch struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"lunch"`
		} `json:"openHour"`
		Disabled bool `json:"disabled"`
	} `json:"organization"`
	Alarmed   bool `json:"alarmed"`
	Agreement struct {
		NoticeReadAt     interface{} `json:"noticeReadAt"`
		AgreedAt         string      `json:"agreedAt"`
		LocationAgreedAt string      `json:"locationAgreedAt"`
		Under14          bool        `json:"under14"`
		Over30           bool        `json:"over30"`
		Over18           bool        `json:"over18"`
	} `json:"agreement"`
	Status    string `json:"status"`
	LeftCount int    `json:"leftCount"`
	Lefts     []struct {
		VaccineType string `json:"vaccineType"`
		VaccineName string `json:"vaccineName"`
		VaccineCode string `json:"vaccineCode"`
		Status      string `json:"status"`
		LeftCount   int    `json:"leftCount"`
	} `json:"lefts"`
	SelectableVaccineCodes []string `json:"selectableVaccineCodes"`
}

type agreementResponse struct {
	NoticeReadAt     string `json:"noticeReadAt"`
	AgreedAt         string `json:"agreedAt"`
	LocationAgreedAt string `json:"locationAgreedAt"`
	Under14          bool   `json:"under14"`
	Over30           bool   `json:"over30"`
	Over18           bool   `json:"over18"`
}

type reservationResponse struct {
	Code         string `json:"code"`
	Desc         string `json:"desc"`
	Organization struct {
		OrgCode     string  `json:"orgCode"`
		OrgName     string  `json:"orgName"`
		ConfirmId   string  `json:"confirmId"`
		PhoneNumber string  `json:"phoneNumber"`
		Address     string  `json:"address"`
		X           float64 `json:"x"`
		Y           float64 `json:"y"`
		OpenHour    struct {
			Date      string `json:"date"`
			DayOfWeek string `json:"dayOfWeek"`
			DayOff    bool   `json:"dayOff"`
			OpenHour  struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"openHour"`
			Lunch struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"lunch"`
		} `json:"openHour"`
		Disabled bool `json:"disabled"`
	} `json:"organization"`
}

type app struct {
	cfg    *Config
	client *http.Client
	tlg    *TelegramLogger
}

func NewApp() (*app, error) {
	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}

	if len(cfg.Locations) == 0 {
		return nil, errors.New("there is no locations to check")
	}

	tlg, err := NewTelegramLogger(cfg.TelegramToken, cfg.TelegramChatID)
	if err != nil {
		return nil, err
	}

	client, err := NewTimeoutClient()
	if err != nil {
		return nil, err
	}

	return &app{
		cfg:    cfg,
		client: client,
		tlg:    tlg,
	}, nil
}

func (s *app) Run() error {
	timePassed := 0
	sleep := s.cfg.Sleep
	notifyAfter := s.cfg.NotifyAfter

	for {
		for _, location := range s.cfg.Locations {
			data := []byte(fmt.Sprintf(`
{
	"bottomRight": {
    	"x": %g,
    	"y": %g
  	},
  	"onlyLeft": true,
  	"order": "latitude",
  	"topLeft": {
		"x": %g,
		"y": %g
  	}
}`, location.BottomRight.X, location.BottomRight.Y, location.TopLeft.X, location.TopLeft.Y))

			log.Printf("data: %s", string(data))
			req, err := newRequest(LEFT_COUNT_URL, data)
			if err != nil {
				log.Printf("failed to create request: %v", err)
				continue
			}

			resp, err := s.client.Do(req)
			if err != nil {
				log.Printf("failed to do request: %v", err)
				continue
			}

			defer resp.Body.Close()

			dat, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("failed to read response body: %v", err)
				continue
			}

			var respJson response
			err = json.Unmarshal(dat, &respJson)
			if err != nil {
				log.Printf("failed to unmarshal json: %v", err)
				continue
			}

			if len(respJson.Organizations) == 0 {
				log.Printf("There is no vaccine left for %s location\n", location.Name)
				continue
			}

			for _, org := range respJson.Organizations {
				timePassed = 0

				//s.tlg.Logf("Hospital: %s\nAddress: %s\nHas: %d vaccine left", org.OrgName, org.Address, org.LeftCounts)
				//s.tlg.Logf("Trying to make reservation")
				//s.tlg.Logf("Getting organization detail information")

				log.Printf("Hospital: %s\nAddress: %s\nHas: %d vaccine left", org.OrgName, org.Address, org.LeftCounts)
				log.Printf("Trying to make reservation")
				log.Printf("Getting organization detail information")

				orgDetail, err := s.org(org.OrgCode)
				if err != nil {
					log.Printf("failed to get organization detail information: %v", err)
					continue
				}

				log.Printf("orgDetail: %+v", orgDetail)

				vaccineCode := ""
				for _, left := range orgDetail.Lefts {
					if left.LeftCount > 0 {
						vaccineCode = left.VaccineCode
						break
					}
				}

				if vaccineCode == "" {
					continue
				}

				log.Printf("vaccineCode: %s", vaccineCode)

				log.Printf("Accept agreement")
				_, err = s.agreement()
				if err != nil {
					log.Printf("failed to accept agreement: %v", err)
					continue
				}

				log.Printf("Make reservation")
				reservation, err := s.reservation(vaccineCode, org.OrgCode)
				if err != nil {
					log.Printf("failed to make reservation: %v", err)
					continue
				}

				log.Printf("Reservation success")
				log.Printf("reservation: %+v", reservation)
				os.Exit(1)
			}
		}

		timePassed += sleep
		log.Printf("Sleep for %d seconds\n", sleep)
		time.Sleep(time.Duration(sleep) * time.Second)

		if timePassed >= notifyAfter {
			s.tlg.Logf("There is still no vaccine")
			timePassed = 0
		}
	}

}

func (s *app) me() (*userResponse, error) {
	req, err := http.NewRequest("GET", ME_URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2")
	req.Header.Set("Cookie", s.cfg.Cookie)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respJson userResponse
	err = json.Unmarshal(dat, &respJson)
	if err != nil {
		return nil, err
	}

	return &respJson, nil
}

func (s *app) org(orgCode string) (*orgResponse, error) {
	req, err := http.NewRequest("GET", ORG_URL+"/"+orgCode, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2")
	req.Header.Set("Cookie", s.cfg.Cookie)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respJson orgResponse
	err = json.Unmarshal(dat, &respJson)
	if err != nil {
		return nil, err
	}

	return &respJson, nil
}

func (s *app) agreement() (*agreementResponse, error) {
	req, err := http.NewRequest("GET", AGREEMENT_URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2")
	req.Header.Set("Cookie", s.cfg.Cookie)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respJson agreementResponse
	err = json.Unmarshal(dat, &respJson)
	if err != nil {
		return nil, err
	}

	return &respJson, nil
}

func (s *app) reservation(vaccineCode string, orgCode string) (*reservationResponse, error) {
	code, _ := strconv.Atoi(orgCode)

	params := fmt.Sprintf(`
{
	"from": "Map",
	"vaccineCode": "%s",
	"orgCode": %d,
	"distance": null
}`, vaccineCode, code)

	log.Printf("params: %s", params)
	buf := bytes.NewBuffer([]byte(params))

	req, err := http.NewRequest("POST", RESERVATION_URL, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2")
	req.Header.Set("Cookie", s.cfg.Cookie)
	req.Header.Set("Content-Length", strconv.Itoa(buf.Len()))
	req.Header.Set("Origin", "https://vaccine-map.kakao.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept", "application/json, text/plain, */*")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respJson reservationResponse
	err = json.Unmarshal(dat, &respJson)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("failed to make reservation, status: %d, response: %s", resp.StatusCode, string(dat)))
	}

	return &respJson, err
}

func parseCookies(cookie string) []*http.Cookie {
	split := strings.Split(cookie, ";")
	var cookies []*http.Cookie

	for _, s := range split {
		nameValue := strings.Split(s, "=")

		name := strings.TrimSpace(nameValue[0])
		value := strings.TrimSpace(nameValue[1])

		log.Printf("name: %s, value: %s\n", name, value)

		cookie := &http.Cookie{
			Name:  name,
			Value: value,
		}

		cookies = append(cookies, cookie)
	}

	return cookies
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
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2")
	req.Header.Set("Referer", "https://vaccine-map.kakao.com/map2?v=1")

	return req, nil
}
