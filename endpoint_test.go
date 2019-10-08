package main

import (
	"encoding/json"
	"testing"

	domain "banners.domain"
)

func TestGetBannersOkResult(t *testing.T) {
	s := GetBanners("UK", "1", "1")

	var results []domain.Zone

	resultsError := json.Unmarshal([]byte(s), &results)

	if resultsError != nil {
		t.Errorf("Error marshaling result: " + resultsError.Error())
	}

	if len(results) == 0 {
		t.Errorf("Error retrieving result: expected values, received empty instead")
	}
}

func TestGetBannersErrorResult(t *testing.T) {
	s := GetBanners("NonExistentLanguageCode", "1", "1")

	var results []domain.Zone

	resultsError := json.Unmarshal([]byte(s), &results)

	if resultsError != nil {
		t.Errorf("Error marshaling result: " + resultsError.Error())
	}

	if len(results) > 0 {
		t.Errorf("Error retrieving result: expected empty, received value instead")
	}
}
