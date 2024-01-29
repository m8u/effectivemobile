package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

type agifyResponse struct {
	Age int8 `json:"age"`
}

func getAge(name string) (int8, error) {
	resp, err := http.Get(os.Getenv("AGIFY_URL") + "/?name=" + name)
	if err != nil {
		log.Debugf("failed to get age for '%s'", name)
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var agifyResp agifyResponse
	err = json.Unmarshal(body, &agifyResp)
	if err != nil {
		return 0, err
	}
	return agifyResp.Age, nil
}

type genderizeResponse struct {
	Gender Gender `json:"gender"`
}

func getGender(name string) (Gender, error) {
	resp, err := http.Get(os.Getenv("GENDERIZE_URL") + "/?name=" + name)
	if err != nil {
		log.Debugf("failed to get gender for '%s'", name)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var genderizeResp genderizeResponse
	err = json.Unmarshal(body, &genderizeResp)
	if err != nil {
		return "", err
	}
	return genderizeResp.Gender, nil
}

type countryEntry struct {
	CountryId Nationality `json:"country_id"`
}
type nationalizeResponse struct {
	Country []countryEntry `json:"country"`
}

func getNationality(name string) (Nationality, error) {
	resp, err := http.Get(os.Getenv("NATIONALIZE_URL") + "/?name=" + name)
	if err != nil {
		log.Debugf("failed to get nationality for '%s'", name)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var nationalizeResp nationalizeResponse
	err = json.Unmarshal(body, &nationalizeResp)
	if err != nil {
		return "", err
	}
	return nationalizeResp.Country[0].CountryId, nil
}
